// Package strings covers string building and JSON handling.
package strings

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang-cookbook/cookbook"
)

// Builder demonstrates strings.Builder for efficient string concatenation.
func Builder() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "strings-builder",
		Category: "strings",
		Title:    "strings.Builder",
		Summary:  "Prefer strings.Builder over += in loops to avoid O(n²) allocations.",
		References: []cookbook.Reference{
			{Title: "pkg.go.dev/strings#Builder", URL: "https://pkg.go.dev/strings#Builder"},
		},
		Run: func() error {
			words := []string{"Go", "is", "fast"}

			var b strings.Builder
			b.Grow(20)
			for i, w := range words {
				if i > 0 {
					b.WriteByte(' ')
				}
				b.WriteString(w)
			}
			fmt.Printf("joined: %q\n", b.String())
			return nil
		},
	}
}

// JSON demonstrates encoding/json for structs and maps.
func JSON() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "strings-json",
		Category: "strings",
		Title:    "JSON Marshal and Unmarshal",
		Summary:  "Use struct tags for field names; prefer json.Decoder for streaming large inputs.",
		References: []cookbook.Reference{
			{Title: "encoding/json package", URL: "https://pkg.go.dev/encoding/json"},
		},
		Run: func() error {
			type Person struct {
				Name  string `json:"name"`
				Age   int    `json:"age"`
				Email string `json:"email,omitempty"`
			}

			p := Person{Name: "Gopher", Age: 10}
			data, err := json.Marshal(p)
			if err != nil {
				return err
			}
			fmt.Printf("marshaled: %s\n", data)

			var decoded Person
			if err := json.Unmarshal(data, &decoded); err != nil {
				return err
			}
			fmt.Printf("decoded: %+v\n", decoded)
			return nil
		},
	}
}
