// Package interfaces covers Go's implicit interface satisfaction and design patterns.
package interfaces

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/golang-cookbook/cookbook"
)

type user struct {
	id   int
	name string
}

type userStore struct {
	users map[int]user
}

func newUserStore() *userStore {
	return &userStore{users: map[int]user{1: {id: 1, name: "Alice"}}}
}

func (s *userStore) get(id int) (user, bool) {
	u, ok := s.users[id]
	return u, ok
}

type userGetter interface {
	get(id int) (user, bool)
}

// SmallInterfaces demonstrates keeping interfaces small (often one method).
func SmallInterfaces() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "interfaces-small",
		Category: "interfaces",
		Title:    "Small, Focused Interfaces",
		Summary:  "Define interfaces at the consumer; io.Reader and io.Writer are the canonical examples.",
		References: []cookbook.Reference{
			{Title: "Effective Go — Interfaces", URL: "https://go.dev/doc/effective_go#interfaces"},
			{Title: "Go Proverbs", URL: "https://go-proverbs.github.io/"},
		},
		Run: func() error {
			// Any type with Write([]byte) (int, error) satisfies io.Writer
			writeUpper := func(w io.Writer, s string) error {
				_, err := fmt.Fprintf(w, "%s", strings.ToUpper(s))
				return err
			}

			var buf bytes.Buffer
			if err := writeUpper(&buf, "hello"); err != nil {
				return err
			}
			fmt.Printf("uppercased via io.Writer: %q\n", buf.String())
			return nil
		},
	}
}

// TypeAssertions demonstrates the comma-ok idiom for safe type assertions.
func TypeAssertions() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "interfaces-type-assertions",
		Category: "interfaces",
		Title:    "Type Assertions and Type Switches",
		Summary:  "Use v, ok := x.(T) for safe assertions; type switches handle multiple concrete types.",
		References: []cookbook.Reference{
			{Title: "A Tour of Go — Type assertions", URL: "https://go.dev/tour/methods/15"},
		},
		Run: func() error {
			var i any = "gopher"

			if s, ok := i.(string); ok {
				fmt.Printf("asserted string: %q\n", s)
			}

			describe := func(v any) {
				switch x := v.(type) {
				case int:
					fmt.Printf("  int: %d\n", x)
				case string:
					fmt.Printf("  string: %q\n", x)
				default:
					fmt.Printf("  unknown: %T\n", x)
				}
			}

			describe(42)
			describe("go")
			describe(3.14)
			return nil
		},
	}
}

// AcceptReturn demonstrates "accept interfaces, return structs".
func AcceptReturn() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "interfaces-accept-return",
		Category: "interfaces",
		Title:    "Accept Interfaces, Return Structs",
		Summary:  "Function parameters should be interfaces; return concrete types so callers can use all methods.",
		References: []cookbook.Reference{
			{Title: "Go Proverbs — Accept interfaces, return structs", URL: "https://go-proverbs.github.io/"},
		},
		Run: func() error {
			greet := func(g userGetter, id int) string {
				u, ok := g.get(id)
				if !ok {
					return "unknown user"
				}
				return "hello, " + u.name
			}

			store := newUserStore()
			fmt.Println(greet(store, 1))
			return nil
		},
	}
}
