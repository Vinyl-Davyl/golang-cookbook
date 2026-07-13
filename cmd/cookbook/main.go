// Command cookbook is the CLI for browsing and running Go Cookbook recipes.
//
// Usage:
//
//	go run ./cmd/cookbook list
//	go run ./cmd/cookbook list errors
//	go run ./cmd/cookbook run errors-wrap
//	go run ./cmd/cookbook doc errors-wrap
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/golang-cookbook/cookbook"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		category := ""
		if len(os.Args) > 2 {
			category = os.Args[2]
		}
		listRecipes(category)
	case "run":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: cookbook run <recipe-id>")
			os.Exit(1)
		}
		runRecipe(os.Args[2])
	case "doc":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: cookbook doc <recipe-id>")
			os.Exit(1)
		}
		showDoc(os.Args[2])
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`Go Cookbook — idiomatic Go recipes with runnable examples

Commands:
  list [category]   List all recipes, optionally filtered by category
  run <recipe-id>   Run a recipe and print its output
  doc <recipe-id>   Show recipe documentation and references

Examples:
  go run ./cmd/cookbook list
  go run ./cmd/cookbook list errors
  go run ./cmd/cookbook run errors-wrap
  go run ./cmd/cookbook doc concurrency-errgroup`)
}

func listRecipes(category string) {
	cats := cookbook.Categories()
	names := make([]string, 0, len(cats))
	for name := range cats {
		if category == "" || name == category {
			names = append(names, name)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("\n## %s\n", name)
		for _, r := range cats[name] {
			fmt.Printf("  %-28s %s\n", r.ID, r.Title)
		}
	}
	fmt.Println()
}

func runRecipe(id string) {
	r, ok := cookbook.ByID(id)
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown recipe %q — run 'list' to see available recipes\n", id)
		os.Exit(1)
	}

	fmt.Printf("=== %s ===\n%s\n\n", r.Title, r.Summary)
	if err := r.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "recipe failed: %v\n", err)
		os.Exit(1)
	}
}

func showDoc(id string) {
	r, ok := cookbook.ByID(id)
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown recipe %q\n", id)
		os.Exit(1)
	}

	var b strings.Builder
	fmt.Fprintf(&b, "ID:       %s\n", r.ID)
	fmt.Fprintf(&b, "Category: %s\n", r.Category)
	fmt.Fprintf(&b, "Title:    %s\n\n", r.Title)
	fmt.Fprintf(&b, "Summary:\n  %s\n", r.Summary)
	if len(r.References) > 0 {
		b.WriteString("\nReferences:\n")
		for _, ref := range r.References {
			fmt.Fprintf(&b, "  - %s\n    %s\n", ref.Title, ref.URL)
		}
	}
	b.WriteString("\nRun: go run ./cmd/cookbook run " + r.ID + "\n")
	fmt.Print(b.String())
}
