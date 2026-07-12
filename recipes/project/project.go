// Package project covers Go modules and recommended project layout.
package project

import (
	"fmt"

	"github.com/golang-cookbook/cookbook"
)

// ModuleLayout explains idiomatic Go project structure.
func ModuleLayout() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "project-module-layout",
		Category: "project",
		Title:    "Modules and Project Layout",
		Summary:  "Start flat with cmd/ and internal/; grow domain packages only when needed; avoid architecture theater.",
		References: []cookbook.Reference{
			{Title: "Go Modules Reference", URL: "https://go.dev/ref/mod"},
			{Title: "golang-standards/project-layout", URL: "https://github.com/golang-standards/project-layout"},
			{Title: "Go Wiki — Package names", URL: "https://go.dev/wiki/PackageNames"},
		},
		Run: func() error {
			layout := `
Recommended layout for this cookbook and small-to-medium projects:

  golang-cookbook/
  ├── go.mod              # module definition
  ├── README.md           # human-readable index
  ├── cmd/
  │   └── cookbook/       # main application entrypoints
  │       └── main.go
  ├── cookbook/           # shared library code (registry, types)
  │   ├── recipe.go
  │   └── registry.go
  └── recipes/            # topic packages (one package per area)
      ├── basics/
      ├── errors/
      └── concurrency/

Principles:
  • cmd/<app>/main.go — thin main, logic in packages
  • internal/ — packages not importable by external modules
  • Flat by default — add packages when you have a real reason
  • go mod init <module-path> to start a new project
  • go work init for local multi-module development
`
			fmt.Print(layout)
			return nil
		},
	}
}
