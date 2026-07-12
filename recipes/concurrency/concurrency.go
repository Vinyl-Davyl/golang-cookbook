// Package concurrency covers goroutines, channels, and synchronization patterns.
package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/golang-cookbook/cookbook"
)

// Goroutines demonstrates lightweight concurrent execution.
func Goroutines() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "concurrency-goroutines",
		Category: "concurrency",
		Title:    "Goroutines",
		Summary:  "Start goroutines with go f(); always ensure every goroutine has a clear exit path.",
		References: []cookbook.Reference{
			{Title: "A Tour of Go — Goroutines", URL: "https://go.dev/tour/concurrency/1"},
			{Title: "Go Concurrency Patterns", URL: "https://go.dev/blog/pipelines"},
		},
		Run: func() error {
			var wg sync.WaitGroup
			for i := 1; i <= 3; i++ {
				wg.Add(1)
				go func(n int) {
					defer wg.Done()
					fmt.Printf("  goroutine %d\n", n)
				}(i)
			}
			wg.Wait()
			fmt.Println("all goroutines finished")
			return nil
		},
	}
}

// Channels demonstrates unbuffered and buffered channels.
func Channels() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "concurrency-channels",
		Category: "concurrency",
		Title:    "Channels",
		Summary:  "Channels communicate between goroutines; close(ch) signals no more sends; range drains until closed.",
		References: []cookbook.Reference{
			{Title: "A Tour of Go — Channels", URL: "https://go.dev/tour/concurrency/2"},
			{Title: "Share Memory By Communicating", URL: "https://go.dev/blog/codelab-share"},
		},
		Run: func() error {
			ch := make(chan int, 2)
			ch <- 1
			ch <- 2
			close(ch)

			for v := range ch {
				fmt.Printf("  received %d\n", v)
			}
			return nil
		},
	}
}

// Select demonstrates multiplexing channel operations.
func Select() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "concurrency-select",
		Category: "concurrency",
		Title:    "select Statement",
		Summary:  "select waits on multiple channel ops; use time.After for timeouts; default makes non-blocking selects.",
		References: []cookbook.Reference{
			{Title: "A Tour of Go — Select", URL: "https://go.dev/tour/concurrency/5"},
		},
		Run: func() error {
			ch := make(chan string, 1)
			ch <- "ready"

			select {
			case msg := <-ch:
				fmt.Printf("got message: %q\n", msg)
			case <-time.After(100 * time.Millisecond):
				fmt.Println("timed out")
			}
			return nil
		},
	}
}

// Semaphore demonstrates bounded concurrency with a buffered channel.
func Semaphore() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "concurrency-semaphore",
		Category: "concurrency",
		Title:    "Bounded Concurrency (Semaphore)",
		Summary:  "Limit parallel work with a buffered channel semaphore instead of a rigid worker pool.",
		References: []cookbook.Reference{
			{Title: "Go Concurrency Patterns — Bounded parallelism", URL: "https://go.dev/blog/pipelines"},
		},
		Run: func() error {
			const limit = 2
			sem := make(chan struct{}, limit)
			var wg sync.WaitGroup

			work := func(id int) {
				defer wg.Done()
				defer func() { <-sem }()
				fmt.Printf("  worker %d running\n", id)
				time.Sleep(50 * time.Millisecond)
			}

			for i := 1; i <= 5; i++ {
				wg.Add(1)
				sem <- struct{}{}
				go work(i)
			}
			wg.Wait()
			fmt.Println("all work done with max 2 concurrent")
			return nil
		},
	}
}

// ProducerConsumer demonstrates the classic channel pattern.
func ProducerConsumer() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "concurrency-producer-consumer",
		Category: "concurrency",
		Title:    "Producer-Consumer Pattern",
		Summary:  "One goroutine produces values on a channel; another consumes until the channel is closed.",
		References: []cookbook.Reference{
			{Title: "Go Concurrency Patterns", URL: "https://go.dev/blog/pipelines"},
		},
		Run: func() error {
			jobs := make(chan int, 5)

			go func() {
				defer close(jobs)
				for i := 1; i <= 5; i++ {
					jobs <- i
				}
			}()

			for j := range jobs {
				fmt.Printf("  processed job %d\n", j)
			}
			return nil
		},
	}
}

// ErrGroup demonstrates golang.org/x/sync/errgroup for concurrent tasks with error propagation.
func ErrGroup() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "concurrency-errgroup",
		Category: "concurrency",
		Title:    "errgroup for Concurrent Tasks",
		Summary:  "errgroup runs goroutines, returns the first error, and cancels sibling tasks via context.",
		References: []cookbook.Reference{
			{Title: "golang.org/x/sync/errgroup", URL: "https://pkg.go.dev/golang.org/x/sync/errgroup"},
		},
		Run: func() error {
			g, ctx := errgroup.WithContext(context.Background())
			g.SetLimit(2)

			tasks := []string{"fetch-users", "fetch-orders", "fetch-products"}
			for _, task := range tasks {
				task := task
				g.Go(func() error {
					select {
					case <-ctx.Done():
						return ctx.Err()
					case <-time.After(30 * time.Millisecond):
						fmt.Printf("  completed %s\n", task)
						return nil
					}
				})
			}

			if err := g.Wait(); err != nil {
				return err
			}
			fmt.Println("all tasks completed")
			return nil
		},
	}
}
