// Package basics covers fundamental Go syntax and control flow.
//
// Start here if you are new to Go. These recipes mirror chapters found in
// "The Go Programming Language" (Donovan & Kernighan) and the official tour.
package basics

import (
	"fmt"

	"github.com/golang-cookbook/cookbook"
)

// Variables demonstrates declaration styles, zero values, and type inference.
func Variables() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "basics-variables",
		Category: "basics",
		Title:    "Variables, Types, and Zero Values",
		Summary:  "Learn short declaration (:=), explicit types, and Go's zero-value semantics.",
		References: []cookbook.Reference{
			{Title: "A Tour of Go — Basics", URL: "https://go.dev/tour/basics/1"},
			{Title: "Effective Go — Names", URL: "https://go.dev/doc/effective_go#names"},
		},
		Run: func() error {
			// Explicit type
			var name string = "gopher"
			var age int

			// Short declaration (most common inside functions)
			language := "Go"
			version := 1.23

			// Zero values: int→0, string→"", bool→false, pointer→nil
			fmt.Printf("name=%q age=%d (zero value)\n", name, age)
			fmt.Printf("language=%q version=%v\n", language, version)

			// Constants
			const maxRetries = 3
			fmt.Printf("maxRetries=%d\n", maxRetries)
			return nil
		},
	}
}

// ControlFlow demonstrates if, for, switch, and range.
func ControlFlow() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "basics-control-flow",
		Category: "basics",
		Title:    "Control Flow: if, for, switch, range",
		Summary:  "Go has one loop keyword (for) and switch without fallthrough by default.",
		References: []cookbook.Reference{
			{Title: "A Tour of Go — Flow control", URL: "https://go.dev/tour/flowcontrol/1"},
		},
		Run: func() error {
			// if with short statement
			if n := 42; n%2 == 0 {
				fmt.Println("42 is even")
			}

			// Classic for loop
			sum := 0
			for i := 1; i <= 5; i++ {
				sum += i
			}
			fmt.Printf("sum 1..5 = %d\n", sum)

			// switch (no fallthrough unless break omitted intentionally)
			day := "Mon"
			switch day {
			case "Sat", "Sun":
				fmt.Println("weekend")
			default:
				fmt.Println("weekday")
			}

			// range over slice
			items := []string{"go", "fmt", "net"}
			for i, v := range items {
				fmt.Printf("  [%d] %s\n", i, v)
			}
			return nil
		},
	}
}

// DeferPanicRecover demonstrates defer, panic, and recover.
func DeferPanicRecover() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "basics-defer-panic-recover",
		Category: "basics",
		Title:    "defer, panic, and recover",
		Summary:  "Use defer for cleanup; reserve panic/recover for truly exceptional cases only.",
		References: []cookbook.Reference{
			{Title: "Effective Go — Defer, panic, and recover", URL: "https://go.dev/doc/effective_go#defer"},
			{Title: "Go Blog — Defer, Panic, and Recover", URL: "https://go.dev/blog/defer-panic-and-recover"},
		},
		Run: func() error {
			fmt.Println("--- defer demo ---")
			defer fmt.Println("  defer: runs last (LIFO)")
			defer fmt.Println("  defer: runs second")
			fmt.Println("  main body runs first")

			fmt.Println("\n--- recover demo ---")
			safeDivide := func(a, b int) (result int, err error) {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("recovered from panic: %v", r)
					}
				}()
				return a / b, nil
			}

			r, err := safeDivide(10, 2)
			fmt.Printf("  10/2 = %d\n", r)

			_, err = safeDivide(10, 0)
			fmt.Printf("  10/0 error: %v\n", err)
			return nil
		},
	}
}
