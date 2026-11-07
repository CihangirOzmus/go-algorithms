package removeelement

import (
	"slices"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want []int // remaining elements, order-independent
	}{
		{"example 1", []int{3, 2, 2, 3}, 3, []int{2, 2}},
		{"example 2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 0, 1, 3, 4}},
		{"empty", []int{}, 1, []int{}},
		{"all removed", []int{1, 1, 1}, 1, []int{}},
		{"none removed", []int{1, 2, 3}, 4, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := RemoveElement(tt.nums, tt.val)
			got := slices.Clone(tt.nums[:k])
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("k=%d, first k = %v, want %v", k, got, tt.want)
			}
		})
	}
}
