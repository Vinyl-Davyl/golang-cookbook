// Package testingrec demonstrates idiomatic Go testing patterns.
package testingrec

import (
	"fmt"
	"testing"

	"github.com/golang-cookbook/cookbook"
)

// TableDriven demonstrates table-driven tests — the standard Go testing style.
func TableDriven() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "testing-table-driven",
		Category: "testing",
		Title:    "Table-Driven Tests",
		Summary:  "Define []struct{ name, input, want } and loop with t.Run for clear, exhaustive cases.",
		References: []cookbook.Reference{
			{Title: "Go Wiki — TableDrivenTests", URL: "https://go.dev/wiki/TableDrivenTests"},
		},
		Run: func() error {
			add := func(a, b int) int { return a + b }

			tests := []struct {
				name string
				a, b int
				want int
			}{
				{"zeros", 0, 0, 0},
				{"positive", 2, 3, 5},
				{"negative", -1, 1, 0},
			}

			for _, tt := range tests {
				got := add(tt.a, tt.b)
				status := "PASS"
				if got != tt.want {
					status = "FAIL"
				}
				fmt.Printf("  %s %s: got %d want %d\n", status, tt.name, got, tt.want)
			}
			return nil
		},
	}
}

// SubtestsAndParallel demonstrates t.Run subtests and t.Parallel.
func SubtestsAndParallel() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "testing-subtests-parallel",
		Category: "testing",
		Title:    "Subtests and Parallel Execution",
		Summary:  "Use t.Run for subtests; t.Parallel() speeds up independent cases; t.Helper() marks test helpers.",
		References: []cookbook.Reference{
			{Title: "testing package", URL: "https://pkg.go.dev/testing"},
		},
		Run: func() error {
			fmt.Println("Run the real tests with:")
			fmt.Println("  go test ./recipes/testingrec/...")
			fmt.Println("")
			fmt.Println("Example pattern:")
			fmt.Println(`  func TestFoo(t *testing.T) {
      t.Parallel()
      for _, tt := range tests {
          tt := tt
          t.Run(tt.name, func(t *testing.T) {
              t.Parallel()
              // assert
          })
      }
  }`)
			return nil
		},
	}
}

func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
