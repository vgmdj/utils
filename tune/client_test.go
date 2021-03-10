package tune

import (
	"context"
	"sync"
	"testing"
	"time"
)

var postMetric = NewSlideWindow(100)

func TestClient_ListenAndResize(t *testing.T) {
	go func() {

		values := []int64{200, 1000, 3000, 5000, 8000, 4000, 10000, 100030}
		for index := 0; ; index++ {
			index = index % len(values)
			for i := 0; i < 1000; i++ {
				postMetric.Observe(values[index])
				time.Sleep(time.Millisecond * 5)
			}

		}
	}()

	cli := NewClient(
		WithResize(func(ctx context.Context, observation, toSize int64) error {
			t.Logf("now %d, resize to %d", observation, toSize)
			return nil
		}),
		WithCollector(postMetric),
		WithArithmeticUnit(NewCongestionTune(&CongestionTuneConfig{
			MinSize:    80,
			MaxSize:    150,
			UpperBound: 10000,
		})),
	)

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		cli.ListenAndResize(ctx)
	}()

	go func() {
		time.Sleep(time.Second * 50)
		cancel()
	}()

	wg.Wait()

}
