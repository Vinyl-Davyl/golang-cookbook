// Package errors collects idiomatic error-handling patterns in Go.
package errors

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/golang-cookbook/cookbook"
)

type temporaryError struct{ retryAfter int }

func (e temporaryError) Error() string {
	return fmt.Sprintf("temporary: retry after %ds", e.retryAfter)
}

type validationError struct {
	field   string
	message string
}

func (e validationError) Error() string {
	return fmt.Sprintf("%s: %s", e.field, e.message)
}

// Wrap demonstrates fmt.Errorf with %w for contextual error wrapping.
func Wrap() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "errors-wrap",
		Category: "errors",
		Title:    "Wrapping Errors with Context",
		Summary:  "Add context with fmt.Errorf(\"...: %w\", err) so callers can still inspect the root cause.",
		References: []cookbook.Reference{
			{Title: "Go Blog — Working with Errors in Go 1.13", URL: "https://go.dev/blog/go1.13-errors"},
		},
		Run: func() error {
			readConfig := func(path string) error {
				_, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("loading config %s: %w", path, err)
				}
				return nil
			}

			err := readConfig("/nonexistent/config.yaml")
			fmt.Printf("wrapped error: %v\n", err)
			fmt.Printf("is not exist? %v\n", errors.Is(err, os.ErrNotExist))
			return nil
		},
	}
}

// IsAndAs demonstrates errors.Is and errors.As for error chain inspection.
func IsAndAs() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "errors-is-as",
		Category: "errors",
		Title:    "errors.Is and errors.As",
		Summary:  "Use Is for sentinel comparison and As to extract typed errors from a wrapped chain.",
		References: []cookbook.Reference{
			{Title: "Go Blog — Working with Errors in Go 1.13", URL: "https://go.dev/blog/go1.13-errors"},
		},
		Run: func() error {
			base := temporaryError{retryAfter: 5}
			wrapped := fmt.Errorf("fetch user: %w", base)

			var temp temporaryError
			if errors.As(wrapped, &temp) {
				fmt.Printf("extracted temporaryError: retry after %ds\n", temp.retryAfter)
			}

			fmt.Printf("errors.Is EOF? %v\n", errors.Is(wrapped, io.EOF))
			return nil
		},
	}
}

// Join demonstrates errors.Join for aggregating multiple errors (Go 1.20+).
func Join() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "errors-join",
		Category: "errors",
		Title:    "Joining Multiple Errors",
		Summary:  "Use errors.Join to combine validation or shutdown errors without third-party libraries.",
		References: []cookbook.Reference{
			{Title: "Go 1.20 Release Notes — errors.Join", URL: "https://go.dev/doc/go1.20#errors"},
		},
		Run: func() error {
			err1 := errors.New("name is required")
			err2 := errors.New("email is invalid")
			combined := errors.Join(err1, err2)

			fmt.Printf("joined error: %v\n", combined)
			fmt.Printf("errors.Is name required? %v\n", errors.Is(combined, err1))
			return nil
		},
	}
}

// Sentinel demonstrates predefined package-level error values.
func Sentinel() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "errors-sentinel",
		Category: "errors",
		Title:    "Sentinel Errors",
		Summary:  "Export var ErrXxx = errors.New(...) for expected conditions callers can compare with errors.Is.",
		References: []cookbook.Reference{
			{Title: "Go Wiki — CodeReviewComments — Error Strings", URL: "https://go.dev/wiki/CodeReviewComments#error-strings"},
		},
		Run: func() error {
			errNotFound := errors.New("not found")

			findUser := func(id int) error {
				if id == 0 {
					return errNotFound
				}
				return nil
			}

			if err := findUser(0); errors.Is(err, errNotFound) {
				fmt.Println("user not found — handle gracefully")
			}
			return nil
		},
	}
}

// CustomType demonstrates structured errors with extra fields.
func CustomType() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "errors-custom-type",
		Category: "errors",
		Title:    "Custom Error Types",
		Summary:  "Define error structs when callers need fields (HTTP status, error codes) via errors.As.",
		References: []cookbook.Reference{
			{Title: "Effective Go — Errors", URL: "https://go.dev/doc/effective_go#errors"},
		},
		Run: func() error {
			validateEmail := func(email string) error {
				if email == "" {
					return validationError{field: "email", message: "required"}
				}
				return nil
			}

			err := validateEmail("")
			var ve validationError
			if errors.As(err, &ve) {
				fmt.Printf("validation failed on %q: %s\n", ve.field, ve.message)
			}
			return nil
		},
	}
}
