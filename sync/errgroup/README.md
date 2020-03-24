# 使用示例

```
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vgmdj/utils/sync/errgroup"
)

func main() {
	defer log.Println("main exiting")

	g := errgroup.WithCancel(context.Background())
	g.Go(One)
	g.Go(Two)
	g.Go(Three)

	err := g.Wait()
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("successfully finish all workers")

}

func One(ctx context.Context) error {
	// do something with ctx
	defer log.Println("one exiting")

	after := time.After(2 * time.Second)
	select {
	case <-ctx.Done():
		log.Println("one Ctx done")
		return fmt.Errorf("one ctx done")
	case <-after:
		log.Println("success one after 2 seconds")
		return nil
	}
}

func Two(ctx context.Context) error {
	defer log.Println("two exiting")
	// do something with ctx

	after := time.After(5 * time.Second)
	select {
	case <-ctx.Done():
		log.Println("two Ctx done")
		return fmt.Errorf("two ctx done")
	case <-after:
		log.Println("two after 5 seconds")
		//return fmt.Errorf("err two after 5 seconds")
		panic("err two after 5 seconds")
	}

}

func Three(ctx context.Context) error {
	defer log.Println("three exiting")
	// do something with ctx

	after := time.After(10 * time.Second)
	select {
	case <-ctx.Done():
		log.Println("three Ctx done")
		return fmt.Errorf("three ctx done")
	case <-after:
		log.Println("three after 10 seconds")
		return fmt.Errorf("err three after 10 seconds")
	}
}




```