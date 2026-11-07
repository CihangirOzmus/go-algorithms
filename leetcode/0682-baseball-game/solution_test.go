package baseballgame

import "testing"

func TestCalPoints(t *testing.T) {
	tests := []struct {
		name string
		ops  []string
		want int
	}{
		{"example 1", []string{"5", "2", "C", "D", "+"}, 30},
		{"example 2", []string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
		{"example 3", []string{"1", "C"}, 0},
		{"only numbers", []string{"1", "2", "3"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalPoints(tt.ops); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
