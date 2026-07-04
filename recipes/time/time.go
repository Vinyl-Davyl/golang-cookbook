// Package time covers timers, tickers, and time formatting.
package time

import (
	"fmt"
	"time"

	"github.com/golang-cookbook/cookbook"
)

// TickersTimers demonstrates time.After, Timer, and Ticker.
func TickersTimers() cookbook.Recipe {
	return cookbook.Recipe{
		ID:       "time-tickers-timers",
		Category: "time",
		Title:    "Timers and Tickers",
		Summary:  "Use time.After for one-shot delays; NewTicker for repeated events; always stop timers/tickers.",
		References: []cookbook.Reference{
			{Title: "time package", URL: "https://pkg.go.dev/time"},
		},
		Run: func() error {
			// One-shot timer
			timer := time.NewTimer(30 * time.Millisecond)
			<-timer.C
			fmt.Println("timer fired once")

			// Ticker for repeated events
			ticker := time.NewTicker(20 * time.Millisecond)
			defer ticker.Stop()

			for i := 0; i < 3; i++ {
				<-ticker.C
				fmt.Printf("  tick %d\n", i+1)
			}

			// Parsing and formatting (always use RFC3339 for APIs)
			now := time.Now().UTC()
			fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
			return nil
		},
	}
}
