package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

func main() {
	var (
		ErrFound                = errors.New("found")
		proofOfWorkNumber int64 = 1337
		probes            int64
		result            int64
	)

	// g := &errgroup.Group{}
	g, ctx := errgroup.WithContext(context.Background())
	workers := 100
	// ctx, cancel := context.WithCancel()
	for i := 0; i < workers; i++ {
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return nil
				default:
					seed := atomic.AddInt64(&probes, 1)
					source := rand.NewSource(seed)

					number := rand.New(source).Int63()
					if number%proofOfWorkNumber == 0 && number != 0 {
						atomic.StoreInt64(&result, number)
						return ErrFound
					}
				}
			}

		})
	}

	g.Wait()

	fmt.Printf("Found %v at %v probes", result, probes)
}
