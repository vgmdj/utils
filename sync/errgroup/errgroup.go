// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errgroup provides synchronization, error propagation, and Context
// cancelation for groups of goroutines working on subtasks of a common task.
package errgroup

import (
	"context"
	"fmt"
	"sync"
)

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
type Group struct {
	workerOnce sync.Once

	ctx    context.Context
	cancel func()

	err     error
	errOnce sync.Once
	wg      sync.WaitGroup
}

// WithContext create a Group.
// given function from Go will receive this context,
// and will not give cancel func
func WithContext(ctx context.Context) *Group {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Group{ctx: ctx}
}

// WithCancel create a new Group and an associated Context derived from ctx.
//
// given function from Go will receive context derived from this ctx,
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
func WithCancel(ctx context.Context) *Group {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(ctx)
	return &Group{ctx: ctx, cancel: cancel}
}

// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *Group) Go(f func(ctx context.Context) error) {
	g.wg.Add(1)
	go g.do(f)
}

// do add recover and ctx funcs
func (g *Group) do(f func(ctx context.Context) error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("errgroup: panic with: %s", r)
		}
		if err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
		g.wg.Done()

	}()

	err = f(g.ctx)

}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
