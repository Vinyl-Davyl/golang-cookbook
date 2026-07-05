// Package syncrec covers sync primitives: Mutex, WaitGroup, and Once.
package syncrec

import (
	"fmt"
	"sync"

	"github.com/golang-cookbook/cookbook"
)

// MutexWaitGroup demonstrates protecting shared state and waiting for goroutines.
func MutexWaitGroup() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "sync-mutex-waitgroup",
		Category: "sync",
		Title:    "sync.Mutex and sync.WaitGroup",
		Summary:  "Protect shared mutable state with Mutex; WaitGroup waits for a group of goroutines to finish.",
		References: []cookbook.Reference{
			{Title: "sync package", URL: "https://pkg.go.dev/sync"},
		},
		Run: func() error {
			var (
				mu    sync.Mutex
				wg    sync.WaitGroup
				count int
			)

			for i := 0; i < 100; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					mu.Lock()
					count++
					mu.Unlock()
				}()
			}
			wg.Wait()
			fmt.Printf("final count: %d\n", count)
			return nil
		},
	}
}

// Once demonstrates sync.Once for one-time initialization.
func Once() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "sync-once",
		Category: "sync",
		Title:    "sync.Once for Lazy Initialization",
		Summary:  "sync.Once guarantees a function runs exactly once, even with concurrent callers.",
		References: []cookbook.Reference{
			{Title: "sync.Once", URL: "https://pkg.go.dev/sync#Once"},
		},
		Run: func() error {
			var once sync.Once
			var config string

			init := func() {
				config = "loaded"
				fmt.Println("  config initialized (once)")
			}

			var wg sync.WaitGroup
			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					once.Do(init)
				}()
			}
			wg.Wait()
			fmt.Printf("config = %q\n", config)
			return nil
		},
	}
}
