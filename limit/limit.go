package limit

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/go-redis/redis"
)

const (
	// DefaultLockTimeout default handler lockTimeout
	DefaultLockTimeout = 30 * time.Second
)

// Rules the array of rule
type Rules []Rule

// MinStart the min start of rules
func (r Rules) MinStart() int64 {
	minStart := time.Date(9999, 12, 31, 23, 59, 59, 0, time.Local)
	for _, c := range r {
		if minStart.After(c.Start) {
			minStart = c.Start
		}
	}

	return minStart.UnixNano() / 1e6
}

// MaxEnd the max end of rules
func (r Rules) MaxEnd() int64 {
	max := time.Time{}
	for _, c := range r {
		if max.Before(c.End) {
			max = c.End
		}
	}

	return max.UnixNano() / 1e6
}

//Rule given rule, time is unix nanosecond
type Rule struct {
	Start time.Time // start time
	End   time.Time // end time
	Count int64     // max amount in [start,end]
}

func isFollowRules(rules Rules, ts []int64) bool {
	m := make([]int64, len(rules))
	for i, rule := range rules {
		for _, c := range ts {
			if c <= rule.End.UnixNano()/1e6 && c >= rule.Start.UnixNano()/1e6 {
				m[i]++
			}
		}

		if m[i] >= rule.Count {
			return false
		}
	}

	return true
}

// Limit interface
type Limit interface {
	IsLimit(ctx context.Context, key string, rules Rules) (bool, error)
	Exec(ctx context.Context, key string, rules Rules, timeout time.Duration) error
}

var _ Limit = (*RuleLimit)(nil)

// RuleLimit limit
type RuleLimit struct {
	conn           *redis.Client
	fn             Handler
	unLockFailedFn UnlockFailedHandler
	lockTimeout    time.Duration
	prefix         string
}

// Handler func type
type Handler func(ctx context.Context) error
type UnlockFailedHandler func(ctx context.Context, errBefore error)

var DefaultUnlockFailedHandler = func(ctx context.Context, errBefore error) {
	if errBefore != nil {
		log.Println(errBefore.Error())
		return
	}
}

// NewRuleLimit return the object of rule limit
func NewRuleLimit(conn *redis.Client, fn Handler, ops ...Options) *RuleLimit {
	l := &RuleLimit{
		conn:           conn,
		fn:             fn,
		unLockFailedFn: DefaultUnlockFailedHandler,
		lockTimeout:    DefaultLockTimeout,
	}

	for _, op := range ops {
		op(l)
	}

	return l

}

// Options limit options
type Options func(limit *RuleLimit)

// WithTimeout use given lockTimeout instead of default one
func WithTimeout(second int64) Options {
	return func(limit *RuleLimit) {
		limit.lockTimeout = time.Duration(second) * time.Second
	}
}

// WithTimeout use given lockTimeout instead of default one
func WithPrefix(prefix string) Options {
	return func(limit *RuleLimit) {
		limit.prefix = prefix
	}
}

// WithUnlockFailed the handler when redis unlock failed
func WithUnlockFailed(handler UnlockFailedHandler) Options {
	return func(limit *RuleLimit) {
		limit.unLockFailedFn = handler
	}
}

// IsLimit check the key is available or not
func (l *RuleLimit) IsLimit(ctx context.Context, key string, rules Rules) (bool, error) {
	ts, err := l.getTimestamps(key, rules.MinStart(), rules.MaxEnd())
	if err != nil {
		return false, &RedisError{redisErr: "get timestamp failed," + err.Error()}
	}

	if !isFollowRules(rules, ts) {
		return false, nil
	}

	return true, nil
}

// Exec exec
func (l *RuleLimit) Exec(ctx context.Context, key string, rules Rules, expiration time.Duration) (err error) {
	randomValue := uuid.New().String()
	err = l.lock(key, randomValue)
	if err != nil {
		return &RedisError{redisErr: err.Error()}
	}
	defer func() {
		ulErr := l.unlock(key, randomValue)
		if ulErr != nil {
			l.unLockFailedFn(ctx, err)
		}
	}()

	ts, err := l.getTimestamps(key, rules.MinStart(), rules.MaxEnd())
	if err != nil {
		return &RedisError{redisErr: "get timestamp failed," + err.Error()}
	}

	if !isFollowRules(rules, ts) {
		return &OverLimitError{limitErr: "Frequency limit exceeded"}
	}

	err = l.fn(ctx)
	if err != nil {
		return &HandlerError{handlerErr: err.Error()}
	}

	err = l.addTimestamp(key, time.Now().UnixNano()/1e6, expiration, rules.MinStart())
	if err != nil {
		return &RedisError{redisErr: err.Error()}
	}

	return nil

}

const (
	actionLock  = "lock"
	actionLimit = "limit"
)

// redisKey the final key in redis
func (l *RuleLimit) redisKey(key string, action string) string {
	return fmt.Sprintf("%s[%s-%s]", l.prefix, key, action)
}

// lock redis SetNx lock
func (l *RuleLimit) lock(key string, value string) error {
	ok, err := l.conn.SetNX(l.redisKey(key, actionLock), value, l.lockTimeout).Result()
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("redis lock failed, because key is already exist")
	}

	return nil
}

var unlockScript = redis.NewScript(`
	if redis.call('get',KEYS[1]) == ARGV[1] then
		return redis.call('del',KEYS[1])
	else
		return 0
	end
`)

// unlock redis unlock
func (l *RuleLimit) unlock(key, value string) error {
	ok, err := unlockScript.Run(l.conn, []string{l.redisKey(key, actionLock)}, value).Bool()
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("redis unlock failed")
	}

	return nil

}

func (l *RuleLimit) getTimestamps(key string, minStart, maxEnd int64) ([]int64, error) {
	res, err := l.conn.ZRangeByScore(l.redisKey(key, actionLimit), redis.ZRangeBy{
		Min:    strconv.FormatInt(minStart, 10),
		Max:    strconv.FormatInt(maxEnd, 10),
		Offset: 0,
		Count:  0,
	}).Result()
	if err != nil {
		return nil, err
	}

	scores := make([]int64, len(res))
	for i, c := range res {
		scores[i], err = strconv.ParseInt(c, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return scores, nil
}

func (l *RuleLimit) addTimestamp(key string, ts int64, expiration time.Duration, minStart int64) error {
	p := l.conn.Pipeline()
	zRem := p.ZRemRangeByScore(l.redisKey(key, actionLimit), "-inf",
		strconv.FormatInt(minStart, 10))
	p.Process(zRem)

	zAddCmd := p.ZAdd(l.redisKey(key, actionLimit), redis.Z{
		Score:  float64(ts),
		Member: ts,
	})
	p.Process(zAddCmd)

	expireCmd := p.Expire(l.redisKey(key, actionLimit), expiration)
	p.Process(expireCmd)

	_, err := p.Exec()
	if err != nil {
		return err
	}

	return nil

}
