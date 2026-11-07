package minstack

import "testing"

func TestMinStackExample(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if got := s.GetMin(); got != -3 {
		t.Fatalf("GetMin = %d, want -3", got)
	}
	s.Pop()
	if got := s.Top(); got != 0 {
		t.Fatalf("Top = %d, want 0", got)
	}
	if got := s.GetMin(); got != -2 {
		t.Fatalf("GetMin = %d, want -2", got)
	}
}

func TestMinStackDuplicateMins(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(1)
	s.Push(2)
	s.Pop()
	s.Pop()
	if got := s.GetMin(); got != 1 {
		t.Fatalf("GetMin = %d, want 1 (one copy of the min is still on the stack)", got)
	}
	s.Push(-5)
	if got := s.GetMin(); got != -5 {
		t.Fatalf("GetMin = %d, want -5", got)
	}
	s.Pop()
	if got := s.GetMin(); got != 1 {
		t.Fatalf("GetMin after popping -5 = %d, want 1", got)
	}
}
