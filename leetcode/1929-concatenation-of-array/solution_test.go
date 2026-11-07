package concatenationofarray

import (
	"slices"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example 1", []int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
		{"example 2", []int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
		{"single", []int{5}, []int{5, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConcatenation(tt.nums); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
