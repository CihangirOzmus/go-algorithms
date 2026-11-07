package replaceelementswithgreatestelementonrightside

import (
	"slices"
	"testing"
)

func TestReplaceElements(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"example 1", []int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
		{"example 2", []int{400}, []int{-1}},
		{"increasing", []int{1, 2, 3}, []int{3, 3, -1}},
		{"decreasing", []int{3, 2, 1}, []int{2, 1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceElements(tt.arr); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
