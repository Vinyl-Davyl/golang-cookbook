// Package cookbook defines the recipe model and registry for the Go Cookbook.
//
// Each recipe is runnable Go code with documentation you can browse via:
//
//   - go doc ./...
//   - go run ./cmd/cookbook list
//   - go run ./cmd/cookbook run <recipe-id>
package cookbook

import "fmt"

// Reference links a recipe to external documentation (official docs, blog posts, books).
type Reference struct {
	Title string
	URL   string
}

// Recipe is a single cookbook entry: documented, categorized, and runnable.
type Recipe struct {
	// ID is a stable kebab-case identifier, e.g. "errors-wrap".
	ID string

	// Category groups related recipes, e.g. "errors", "concurrency".
	Category string

	// Title is a short human-readable name.
	Title string

	// Summary explains what the recipe teaches and when to use it.
	Summary string

	// References point to authoritative sources (Effective Go, go.dev blog, etc.).
	References []Reference

	// Run executes the recipe and prints output to demonstrate the pattern.
	Run func() error
}

// String returns a one-line description suitable for listing.
func (r Recipe) String() string {
	return fmt.Sprintf("[%s] %s — %s", r.ID, r.Title, r.Summary)
}
