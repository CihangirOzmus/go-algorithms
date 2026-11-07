package reverselinkedlist

import (
	"slices"
	"testing"
)

func build(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *Node) []int {
	var out []int
	for ; head != nil; head = head.Next {
		out = append(out, head.Val)
	}
	return out
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"example 1", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"example 2", []int{1, 2}, []int{2, 1}},
		{"empty", nil, nil},
		{"single", []int{7}, []int{7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toSlice(ReverseList(build(tt.in))); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
