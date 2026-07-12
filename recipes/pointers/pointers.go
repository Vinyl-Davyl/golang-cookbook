// Package pointers explains when to use pointers vs values in Go.
package pointers

import (
	"fmt"

	"github.com/golang-cookbook/cookbook"
)

// WhenToUse demonstrates pointer vs value trade-offs.
func WhenToUse() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "pointers-when-to-use",
		Category: "pointers",
		Title:    "When to Use Pointers",
		Summary:  "Use pointers to mutate callers' data, avoid copying large structs, or represent optional/missing values.",
		References: []cookbook.Reference{
			{Title: "Go Wiki — CodeReviewComments — Receiver Type", URL: "https://go.dev/wiki/CodeReviewComments#receiver-type"},
		},
		Run: func() error {
			type Config struct {
				Host string
				Port int
			}

			// Pointer: mutates the original
			setPort := func(c *Config, port int) {
				c.Port = port
			}

			// Value: works on a copy
			setHostCopy := func(c Config, host string) Config {
				c.Host = host
				return c
			}

			cfg := Config{Host: "localhost", Port: 8080}
			setPort(&cfg, 9090)
			fmt.Printf("after setPort: %+v\n", cfg)

			cfg = setHostCopy(cfg, "127.0.0.1")
			fmt.Printf("after setHostCopy: %+v\n", cfg)

			// nil pointer = "no value" for optional fields
			var optional *string
			if optional == nil {
				fmt.Println("optional string is nil (not set)")
			}
			return nil
		},
	}
}
