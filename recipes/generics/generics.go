// Package generics covers Go 1.18+ generic functions and type constraints.
package generics

import (
	"cmp"
	"fmt"

	"github.com/golang-cookbook/cookbook"
)

// Functions demonstrates generic functions with type parameters.
func Functions() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "generics-functions",
		Category: "generics",
		Title:    "Generic Functions",
		Summary:  "Use type parameters when the same algorithm works across types; don't over-genericize.",
		References: []cookbook.Reference{
			{Title: "Go Tutorial — Generics", URL: "https://go.dev/doc/tutorial/generics"},
			{Title: "Go 1.18 Release Notes", URL: "https://go.dev/doc/go1.18"},
		},
		Run: func() error {
			min := func[T cmp.Ordered](a, b T) T {
				if a < b {
					return a
				}
				return b
			}

			fmt.Printf("min(3, 7) = %d\n", min(3, 7))
			fmt.Printf("min(3.5, 2.1) = %.1f\n", min(3.5, 2.1))
			fmt.Printf("min(\"go\", \"rust\") = %q\n", min("go", "rust"))
			return nil
		},
	}
}

// Constraints demonstrates custom type constraints with interfaces.
func Constraints() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "generics-constraints",
		Category: "generics",
		Title:    "Type Constraints",
		Summary:  "Define constraint interfaces with a union of types or embedded methods.",
		References: []cookbook.Reference{
			{Title: "Type Parameters Proposal", URL: "https://go.dev/design/43651-type-parameters"},
		},
		Run: func() error {
			type Stringer interface {
				~string
			}

			repeat := func[T Stringer](s T, n int) T {
				var result T
				for i := 0; i < n; i++ {
					result += s
				}
				return result
			}

			fmt.Printf("repeat: %q\n", repeat("go", 3))
			return nil
		},
	}
}
