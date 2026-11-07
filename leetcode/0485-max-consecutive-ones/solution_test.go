package maxconsecutiveones

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{1, 1, 0, 1, 1, 1}, 3},
		{"example 2", []int{1, 0, 1, 1, 0, 1}, 2},
		{"all zeros", []int{0, 0, 0}, 0},
		{"all ones", []int{1, 1, 1, 1}, 4},
		{"longest at start", []int{1, 1, 1, 0, 1}, 3},
		{"single one", []int{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxConsecutiveOnes(tt.nums); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
