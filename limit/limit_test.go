package limit

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/go-redis/redis"
)

var redisCli *redis.Client

func initRedis() {
	redisCli = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

}

func successHandler(ctx context.Context) error {

	return nil
}

func BenchmarkRuleLimit(b *testing.B) {
	initRedis()
	defer redisCli.Close()

	b.ResetTimer()

	var limit Limit
	limit = NewRuleLimit(redisCli, successHandler)
	for i := 0; i < b.N; i++ {
		err := limit.Exec(context.Background(), uuid.New().String(), []Rule{
			{
				Start: time.Date(2021, 3, 5, 15, 37, 0, 0, time.Local),
				End:   time.Date(2021, 3, 6, 18, 30, 0, 0, time.Local),
				Count: 10000,
			},
			{
				Start: time.Date(2021, 3, 5, 15, 50, 0, 0, time.Local),
				End:   time.Date(2021, 3, 5, 16, 00, 0, 0, time.Local),
				Count: 1000,
			},
		}, 24*time.Hour)
		if err != nil {
			b.Error(err.Error())
		}

	}

}

func TestExec(t *testing.T) {
	initRedis()
	defer redisCli.Close()

	st := time.Now()
	wg := sync.WaitGroup{}

	var limit Limit
	limit = NewRuleLimit(redisCli, successHandler)

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := limit.Exec(context.Background(), uuid.New().String(), []Rule{
				{
					Start: time.Date(2021, 3, 5, 15, 37, 0, 0, time.Local),
					End:   time.Date(2021, 3, 6, 18, 30, 0, 0, time.Local),
					Count: 10000,
				},
				{
					Start: time.Date(2021, 3, 5, 15, 50, 0, 0, time.Local),
					End:   time.Date(2021, 3, 5, 16, 00, 0, 0, time.Local),
					Count: 1000,
				},
			}, 24*time.Hour)
			if err != nil {
				t.Error(err.Error())
			}
		}()

	}

	wg.Wait()
	et := time.Now()

	t.Log("共耗时", et.Sub(st))

}
