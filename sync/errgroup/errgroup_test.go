package errgroup_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/vgmdj/utils/sync/errgroup"
	"go.uber.org/atomic"
)

type fake struct {
	mutex sync.Mutex
	A     *atomic.Int32
	B     *atomic.Int32
}

func (f *fake) BAdd() int32 {
	f.B.Add(1)
	return f.B.Load()
}

func (f *fake) fakeTask(ctx context.Context) error {
	if f.A.Load() == 100 {
		return errors.New("test fake error")
	}

	f.A.Add(1)

	return nil

}

func (f *fake) delayTest(ctx context.Context) error {
	timeout := time.Duration(500 * time.Millisecond)
	if deadline, ok := ctx.Deadline(); ok {
		if ctimeout := time.Until(deadline); ctimeout < timeout {
			timeout = ctimeout
		}
	}

	select {
	case <-time.After(timeout):
		return errors.New("timeout")

	case <-ctx.Done():
		return nil

	case <-time.After((200-50)*time.Millisecond + time.Duration(f.BAdd())*time.Millisecond):
		f.A.Add(1)
		return nil
	}

}

func TestWithCancel(t *testing.T) {
	g := errgroup.WithCancel(context.Background())
	f := fake{
		A: atomic.NewInt32(1),
		B: atomic.NewInt32(0),
	}

	for i := 0; i < 99; i++ {
		g.Go(f.fakeTask)
	}

	err := g.Wait()
	if err != nil {
		t.Error(err.Error())
		return
	}

	if f.A.Load() != 100 {
		t.Errorf("expected f.A is 100 ,but got %d\n", f.A.Load())
		return
	}

}

func TestDeadline(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(200*time.Millisecond))
	defer cancel()
	g := errgroup.WithCancel(ctx)

	f := fake{
		A: atomic.NewInt32(1),
		B: atomic.NewInt32(0),
	}

	for i := 0; i < 99; i++ {
		g.Go(f.delayTest)
	}

	err := g.Wait()
	if err != nil && err.Error() != "timeout" {
		t.Error(err.Error())
		return
	}

	if f.A.Load() != 50 {
		t.Errorf("expected f.A is 50 ,but got %d\n", f.A.Load())
		return
	}
}
