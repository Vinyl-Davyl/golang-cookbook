// Package slices covers slice and map idioms in Go.
package slices

import (
	"fmt"
	"slices"

	"github.com/golang-cookbook/cookbook"
)

// AppendAndCapacity demonstrates append, len, cap, and slice sharing.
func AppendAndCapacity() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "slices-append-capacity",
		Category: "slices",
		Title:    "Append, Length, and Capacity",
		Summary:  "append may reallocate; subslices share backing arrays — copy when independence matters.",
		References: []cookbook.Reference{
			{Title: "Go Slices: usage and internals", URL: "https://go.dev/blog/slices-intro"},
			{Title: "slices package (Go 1.21+)", URL: "https://pkg.go.dev/slices"},
		},
		Run: func() error {
			s := make([]int, 0, 4)
			for _, v := range []int{1, 2, 3} {
				s = append(s, v)
			}
			fmt.Printf("len=%d cap=%d data=%v\n", len(s), cap(s), s)

			sub := s[1:3]
			sub[0] = 99
			fmt.Printf("after mutating sub[0]: original=%v (shared backing array)\n", s)

			independent := make([]int, len(sub))
			copy(independent, sub)
			independent[0] = 42
			fmt.Printf("independent copy does not affect original: %v\n", s)
			return nil
		},
	}
}

// FilterMap demonstrates functional-style slice transforms without external libs.
func FilterMap() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "slices-filter-map",
		Category: "slices",
		Title:    "Filter and Map Patterns",
		Summary:  "Pre-allocate with make([]T, 0, len) when size is known; use slices package helpers in Go 1.21+.",
		References: []cookbook.Reference{
			{Title: "pkg.go.dev/slices", URL: "https://pkg.go.dev/slices"},
		},
		Run: func() error {
			nums := []int{1, 2, 3, 4, 5, 6}

			evens := make([]int, 0, len(nums))
			for _, n := range nums {
				if n%2 == 0 {
					evens = append(evens, n)
				}
			}
			fmt.Printf("evens: %v\n", evens)

			doubled := make([]int, len(nums))
			for i, n := range nums {
				doubled[i] = n * 2
			}
			fmt.Printf("doubled: %v\n", doubled)

			if slices.Contains(nums, 3) {
				fmt.Println("contains 3")
			}
			return nil
		},
	}
}

// MapPatterns demonstrates map creation, existence checks, and deletion.
func MapPatterns() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "maps-patterns",
		Category: "slices",
		Title:    "Map Idioms",
		Summary:  "Use comma-ok for existence checks; delete(key) removes entries; zero value is nil.",
		References: []cookbook.Reference{
			{Title: "A Tour of Go — Maps", URL: "https://go.dev/tour/moretypes/19"},
		},
		Run: func() error {
			scores := map[string]int{
				"alice": 95,
				"bob":   87,
			}

			if score, ok := scores["alice"]; ok {
				fmt.Printf("alice: %d\n", score)
			}

			if _, ok := scores["carol"]; !ok {
				fmt.Println("carol not found")
			}

			scores["carol"] = 91
			delete(scores, "bob")
			fmt.Printf("final map: %v\n", scores)
			return nil
		},
	}
}
