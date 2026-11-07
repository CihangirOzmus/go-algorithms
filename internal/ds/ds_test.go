package ds

import (
	"slices"
	"testing"
)

func TestList(t *testing.T) {
	want := []int{1, 2, 3}
	if got := ListToSlice(NewList(want...)); !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	if NewList() != nil {
		t.Fatal("empty list should be nil")
	}
}

func TestTree(t *testing.T) {
	want := []int{1, Null, 2, 3}
	if got := TreeToSlice(NewTree(want...)); !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	if NewTree() != nil {
		t.Fatal("empty tree should be nil")
	}
}
