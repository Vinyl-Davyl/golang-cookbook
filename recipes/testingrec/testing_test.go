package testingrec

import "testing"

func add(a, b int) int { return a + b }

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		want     int
		parallel bool
	}{
		{"zeros", 0, 0, 0, true},
		{"positive", 2, 3, 5, true},
		{"negative", -1, 1, 0, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.parallel {
				t.Parallel()
			}
			assertEqual(t, add(tt.a, tt.b), tt.want)
		})
	}
}
