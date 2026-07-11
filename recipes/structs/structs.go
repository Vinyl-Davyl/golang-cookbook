// Package structs covers struct composition and method design.
package structs

import (
	"fmt"

	"github.com/golang-cookbook/cookbook"
)

type engine struct{ hp int }

func (e engine) rev() string { return fmt.Sprintf("revving %d HP", e.hp) }

type car struct {
	engine        // embedded — car gets engine's methods
	model  string
}

type counter struct{ n int }

func (c *counter) inc() { c.n++ }
func (c counter) value() int { return c.n }

// Embedding demonstrates struct embedding for composition over inheritance.
func Embedding() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "structs-embedding",
		Category: "structs",
		Title:    "Struct Embedding (Composition)",
		Summary:  "Embed types to promote their methods; Go favors composition over class inheritance.",
		References: []cookbook.Reference{
			{Title: "Effective Go — Embedding", URL: "https://go.dev/doc/effective_go#embedding"},
		},
		Run: func() error {
			c := car{engine: engine{hp: 200}, model: "Gopher GT"}
			fmt.Println(c.rev()) // promoted method
			fmt.Printf("model: %s\n", c.model)
			return nil
		},
	}
}

// MethodReceivers compares value vs pointer receivers.
func MethodReceivers() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "structs-method-receivers",
		Category: "structs",
		Title:    "Value vs Pointer Receivers",
		Summary:  "Use pointer receivers when methods mutate state or the struct is large; be consistent within a type.",
		References: []cookbook.Reference{
			{Title: "Effective Go — Methods", URL: "https://go.dev/doc/effective_go#methods"},
			{Title: "Go FAQ — Methods on pointers vs values", URL: "https://go.dev/doc/faq#methods_on_pointers_or_values"},
		},
		Run: func() error {
			var c counter
			c.inc()
			c.inc()
			fmt.Printf("counter value: %d\n", c.value())
			return nil
		},
	}
}
