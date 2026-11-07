package findtargetindicesaftersortingarray

import (
	"slices"
	"testing"
)

func TestTargetIndices(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"example 1", []int{1, 2, 5, 2, 3}, 2, []int{1, 2}},
		{"example 2", []int{1, 2, 5, 2, 3}, 3, []int{3}},
		{"example 3", []int{1, 2, 5, 2, 3}, 5, []int{4}},
		{"not found", []int{1, 2, 5, 2, 3}, 4, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TargetIndices(tt.nums, tt.target); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
