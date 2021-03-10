package tune

import (
	"math"
	"sync"
)

// Tune tune interface
type Tune interface {
	SuitableSize(observation int64) int64
}

// assert
var _ Tune = (*CongestionTune)(nil)

// CongestionTuneConfig the config define
type CongestionTuneConfig struct {
	MinSize    int64
	MaxSize    int64
	UpperBound int64
}

// observationQueue the queue of observation
type observationQueue struct {
	mtx    sync.Mutex
	cap    int
	values []int64
}

func newObservationQueue(cap int) *observationQueue {
	if cap <= 0 {
		cap = 100
	}

	return &observationQueue{
		cap:    cap,
		values: make([]int64, 0, cap+1),
	}
}

func (oq *observationQueue) push(v int64) {
	oq.mtx.Lock()
	defer oq.mtx.Unlock()
	oq.values = append(oq.values, v)

	if len(oq.values) > oq.cap {
		oq.values = oq.values[1:]
	}

}

func (oq *observationQueue) list() []int64 {
	oq.mtx.Lock()
	defer oq.mtx.Unlock()

	cp := make([]int64, len(oq.values))
	copy(cp, oq.values)

	return cp

}

// CongestionTune the congestion tune
type CongestionTune struct {
	lastObservation *observationQueue
	recommendSize   int64
	currentLevel    level
	ssthresh        int64
	minPoolSize     int64
	maxPoolSize     int64
	upperBound      int64
}

// NewCongestionTune create a new congestion tune configured with the given config
func NewCongestionTune(c *CongestionTuneConfig) *CongestionTune {
	return &CongestionTune{
		lastObservation: newObservationQueue(100),
		currentLevel:    levelNormal,
		recommendSize:   0,
		ssthresh:        (c.MinSize + c.MaxSize) / 2,
		minPoolSize:     c.MinSize,
		maxPoolSize:     c.MaxSize,
		upperBound:      c.UpperBound,
	}
}

type level string

const (
	levelNormal   level = "normal"
	levelWarn     level = "warning"
	levelCritical level = "critical"
	levelError    level = "error"
)

func (ct *CongestionTune) getLevel(now int64) level {
	bound := []float64{
		float64(ct.upperBound) * 0.5,
		float64(ct.upperBound) * 0.8,
		float64(ct.upperBound),
	}
	levels := []level{levelNormal, levelWarn, levelCritical, levelError}

	for i, c := range bound {
		if float64(now) < c {
			return levels[i]
		}
	}

	return levelError
}

// SuitableSize figure out the suitable size according to observation
func (ct *CongestionTune) SuitableSize(observation int64) int64 {
	ct.lastObservation.push(observation)
	nowLevel := ct.getLevel(observation)
	ct.currentLevel = nowLevel

	interval := ct.maxPoolSize - ct.minPoolSize

	switch nowLevel {
	case levelNormal:
		ct.recommendSize += int64(math.Ceil(float64(interval / 10)))

	case levelWarn:
		ct.recommendSize += int64(math.Ceil(float64(interval / 100)))

	case levelCritical:
		ct.recommendSize -= int64(math.Ceil(float64(interval / 5)))

	case levelError:
		ct.recommendSize -= int64(math.Ceil(float64(interval / 2)))

	}

	if ct.recommendSize < ct.minPoolSize {
		ct.recommendSize = ct.minPoolSize
	}

	if ct.recommendSize > ct.maxPoolSize {
		ct.recommendSize = ct.maxPoolSize
	}

	return ct.recommendSize

}
