// Package functions covers function design idioms in Go.
package functions

import (
	"fmt"

	"github.com/golang-cookbook/cookbook"
)

// MultipleReturns demonstrates Go's idiomatic (value, error) return pattern.
func MultipleReturns() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "functions-multiple-returns",
		Category: "functions",
		Title:    "Multiple Return Values",
		Summary:  "Return (T, error) for fallible operations; check errors immediately at the call site.",
		References: []cookbook.Reference{
			{Title: "Effective Go — Multiple returns", URL: "https://go.dev/doc/effective_go#multiple-returns"},
		},
		Run: func() error {
			sqrt := func(x float64) (float64, error) {
				if x < 0 {
					return 0, fmt.Errorf("sqrt: negative input %v", x)
				}
				// Newton's method (simplified)
				z := x
				for i := 0; i < 10; i++ {
					z -= (z*z - x) / (2 * z)
				}
				return z, nil
			}

			if v, err := sqrt(16); err != nil {
				return err
			} else {
				fmt.Printf("sqrt(16) = %.4f\n", v)
			}

			if _, err := sqrt(-1); err != nil {
				fmt.Printf("sqrt(-1) error: %v\n", err)
			}
			return nil
		},
	}
}

// Variadic demonstrates functions that accept a variable number of arguments.
func Variadic() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "functions-variadic",
		Category: "functions",
		Title:    "Variadic Functions",
		Summary:  "Use ...T to accept zero or more arguments; the parameter becomes a slice inside the function.",
		References: []cookbook.Reference{
			{Title: "A Tour of Go — Variadic Functions", URL: "https://go.dev/tour/moretypes/14"},
		},
		Run: func() error {
			sum := func(nums ...int) int {
				total := 0
				for _, n := range nums {
					total += n
				}
				return total
			}

			fmt.Printf("sum() = %d\n", sum())
			fmt.Printf("sum(1,2,3) = %d\n", sum(1, 2, 3))

			more := []int{4, 5, 6}
			fmt.Printf("sum(more...) = %d\n", sum(more...))
			return nil
		},
	}
}

// Closures demonstrates functions that capture variables from their enclosing scope.
func Closures() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "functions-closures",
		Category: "functions",
		Title:    "Closures and Function Values",
		Summary:  "Functions are first-class values; closures capture variables by reference.",
		References: []cookbook.Reference{
			{Title: "A Tour of Go — Function closures", URL: "https://go.dev/tour/moretypes/25"},
		},
		Run: func() error {
			counter := func() func() int {
				n := 0
				return func() int {
					n++
					return n
				}
			}()

			fmt.Printf("counter() = %d\n", counter())
			fmt.Printf("counter() = %d\n", counter())
			fmt.Printf("counter() = %d\n", counter())
			return nil
		},
	}
}
