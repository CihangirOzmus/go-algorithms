package findthedistancevaluebetweentwoarrays

import "testing"

func TestFindTheDistanceValue(t *testing.T) {
	tests := []struct {
		name       string
		arr1, arr2 []int
		d          int
		want       int
	}{
		{"example 1", []int{4, 5, 8}, []int{10, 9, 1, 8}, 2, 2},
		{"example 2", []int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3, 2},
		{"example 3", []int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6, 1},
		{"d zero, equal values", []int{1, 2}, []int{2}, 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTheDistanceValue(tt.arr1, tt.arr2, tt.d); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
