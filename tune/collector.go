package tune

import (
	"sort"
	"sync"
)

// Observer is the interface that wraps the Observe method, which is used to add observations.
type Observer interface {
	Observe(n int64)
}

// Collector is the interface which is used to get observation
type Collector interface {
	Max() int64
	Min() int64
	Middle() float64
	Avg() float64
	Percent(start, end int64) float64
}

const (
	// DefaultCap default window size
	DefaultCap = 100
	// defaultIdx default index
	defaultIdx = 0
)

// SlideWindow is the collector and observer impl
type SlideWindow struct {
	mutex   sync.Mutex
	cap     int64
	count   int64
	idx     int64
	buckets []int64
}

// NewSlideWindow return the object
func NewSlideWindow(cap int64) *SlideWindow {
	if cap <= 0 {
		cap = DefaultCap
	}

	return &SlideWindow{
		mutex:   sync.Mutex{},
		cap:     cap,
		idx:     defaultIdx,
		count:   defaultIdx,
		buckets: make([]int64, cap),
	}

}

// Observe save the last [cap] value
func (h *SlideWindow) Observe(value int64) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	index := h.idx % h.cap
	h.buckets[index] = value
	h.idx++
	h.count++

}

// getCopySortedBuckets return the copy of sorted bucket
func (h *SlideWindow) getCopySortedBuckets() []int64 {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	count := h.count
	if h.count > h.cap {
		count = h.cap
	}

	newBuckets := make([]int64, count)
	copy(newBuckets, h.buckets)

	sort.Slice(newBuckets, func(i, j int) bool {
		return newBuckets[i] < newBuckets[j]
	})

	return newBuckets
}

// Max the max value of bucket
func (h *SlideWindow) Max() int64 {
	c := h.getCopySortedBuckets()
	if len(c) == 0 {
		return 0
	}

	return c[len(c)-1]
}

// Min the min value of bucket
func (h *SlideWindow) Min() int64 {
	c := h.getCopySortedBuckets()
	if len(c) == 0 {
		return 0
	}
	return c[0]
}

// Middle the middle value of bucket
func (h *SlideWindow) Middle() float64 {
	c := h.getCopySortedBuckets()
	if len(c) == 0 {
		return 0
	}

	if h.cap&1 == 1 {
		return float64(c[h.cap>>1])
	}

	left, right := h.cap>>1, h.cap>>1-1
	return float64((c[left] + c[right]) / 2)

}

// Avg the avg value of bucket
func (h *SlideWindow) Avg() float64 {
	c := h.getCopySortedBuckets()
	if len(c) == 0 {
		return 0
	}

	sum := int64(0)
	for _, v := range c {
		sum += v
	}

	return float64(sum / int64(len(c)))

}

// Percent the num of [start,end], return num/cap
func (h *SlideWindow) Percent(start, end int64) float64 {
	c := h.getCopySortedBuckets()
	if len(c) == 0 {
		return 0
	}

	count := 0
	for _, v := range c {
		if v >= start && v <= end {
			count++
		}
	}

	return float64(count / len(c))

}
