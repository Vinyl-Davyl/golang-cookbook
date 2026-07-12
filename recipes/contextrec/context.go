// Package contextrec covers context.Context patterns for cancellation and deadlines.
package contextrec

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-cookbook/cookbook"
)

// TimeoutCancel demonstrates context.WithTimeout and WithCancel.
func TimeoutCancel() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "context-timeout-cancel",
		Category: "context",
		Title:    "Timeouts and Cancellation",
		Summary:  "Always call cancel() after WithTimeout/WithCancel (defer cancel()); check ctx.Done() in long operations.",
		References: []cookbook.Reference{
			{Title: "context package", URL: "https://pkg.go.dev/context"},
			{Title: "Go Blog — Context", URL: "https://go.dev/blog/context"},
		},
		Run: func() error {
			// Timeout
			ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
			defer cancel()

			select {
			case <-time.After(200 * time.Millisecond):
				fmt.Println("work finished")
			case <-ctx.Done():
				fmt.Printf("timed out: %v\n", ctx.Err())
			}

			// Manual cancel
			ctx2, cancel2 := context.WithCancel(context.Background())
			go func() {
				time.Sleep(20 * time.Millisecond)
				cancel2()
			}()

			<-ctx2.Done()
			fmt.Printf("manually cancelled: %v\n", ctx2.Err())
			return nil
		},
	}
}

// Propagation demonstrates passing context as the first function parameter.
func Propagation() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "context-propagation",
		Category: "context",
		Title:    "Context Propagation",
		Summary:  "Pass ctx context.Context as the first parameter; never store context in structs.",
		References: []cookbook.Reference{
			{Title: "Go Wiki — CodeReviewComments — Contexts", URL: "https://go.dev/wiki/CodeReviewComments#contexts"},
		},
		Run: func() error {
			fetch := func(ctx context.Context, id int) (string, error) {
				select {
				case <-ctx.Done():
					return "", ctx.Err()
				case <-time.After(10 * time.Millisecond):
					return fmt.Sprintf("data-%d", id), nil
				}
			}

			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			data, err := fetch(ctx, 42)
			if err != nil {
				return err
			}
			fmt.Printf("fetched: %s\n", data)
			return nil
		},
	}
}
