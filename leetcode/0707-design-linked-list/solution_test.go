package designlinkedlist

import (
	"slices"
	"testing"
)

type linkedList interface {
	Get(index int) int
	AddAtHead(val int)
	AddAtTail(val int)
	AddAtIndex(index int, val int)
	DeleteAtIndex(index int)
}

func implementations() map[string]func() linkedList {
	return map[string]func() linkedList{
		"singly":         func() linkedList { l := NewSinglyLinkedList(); return &l },
		"singly optimal": func() linkedList { l := NewOptimalSinglyLinkedList(); return &l },
		"doubly":         func() linkedList { l := Constructor(); return &l },
	}
}

// contents reads the list via Get until it returns -1.
func contents(l linkedList) []int {
	var out []int
	for i := 0; ; i++ {
		v := l.Get(i)
		if v == -1 {
			return out
		}
		out = append(out, v)
	}
}

func TestLeetCodeExample(t *testing.T) {
	for name, newList := range implementations() {
		t.Run(name, func(t *testing.T) {
			l := newList()
			l.AddAtHead(1)
			l.AddAtTail(3)
			l.AddAtIndex(1, 2) // 1 2 3
			if got := l.Get(1); got != 2 {
				t.Fatalf("Get(1) = %d, want 2", got)
			}
			l.DeleteAtIndex(1) // 1 3
			if got := l.Get(1); got != 3 {
				t.Fatalf("Get(1) = %d, want 3", got)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	for name, newList := range implementations() {
		t.Run(name, func(t *testing.T) {
			l := newList()
			if got := l.Get(0); got != -1 {
				t.Fatalf("Get on empty = %d, want -1", got)
			}
			l.DeleteAtIndex(0)  // no-op on empty
			l.AddAtIndex(1, 9)  // index > size: ignored
			l.AddAtIndex(0, 10) // 10
			l.AddAtTail(30)     // 10 30
			l.AddAtIndex(2, 40) // index == size appends: 10 30 40
			l.AddAtIndex(1, 20) // 10 20 30 40
			l.AddAtHead(5)      // 5 10 20 30 40
			if got, want := contents(l), []int{5, 10, 20, 30, 40}; !slices.Equal(got, want) {
				t.Fatalf("after adds = %v, want %v", got, want)
			}
			l.DeleteAtIndex(4) // last
			l.DeleteAtIndex(0) // first
			l.DeleteAtIndex(9) // out of range: ignored
			if got, want := contents(l), []int{10, 20, 30}; !slices.Equal(got, want) {
				t.Fatalf("after deletes = %v, want %v", got, want)
			}
		})
	}
}
