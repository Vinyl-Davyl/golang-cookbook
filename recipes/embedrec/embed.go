// Package embedrec demonstrates embedding static files at compile time.
package embedrec

import (
	_ "embed"
	"fmt"

	"github.com/golang-cookbook/cookbook"
)

//go:embed sample.txt
var sampleText string

// EmbedFiles demonstrates the embed package for compile-time file inclusion.
func EmbedFiles() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "embed-files",
		Category: "embed",
		Title:    "Embedding Files with embed",
		Summary:  "Use //go:embed to bundle templates, configs, or migrations into the binary.",
		References: []cookbook.Reference{
			{Title: "embed package", URL: "https://pkg.go.dev/embed"},
			{Title: "Go 1.16 Release Notes — embed", URL: "https://go.dev/doc/go1.16#embed"},
		},
		Run: func() error {
			fmt.Printf("embedded file contents: %q\n", sampleText)
			return nil
		},
	}
}
