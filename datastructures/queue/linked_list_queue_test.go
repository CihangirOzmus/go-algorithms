package queue

import "testing"

func TestEmptyQueue(t *testing.T) {
	var q LinkedListQueue
	if !q.IsEmpty() || q.Size() != 0 {
		t.Fatalf("new queue: IsEmpty=%v Size=%d, want true 0", q.IsEmpty(), q.Size())
	}
	if v, ok := q.Dequeue(); ok {
		t.Fatalf("Dequeue on empty = (%d, true), want ok=false", v)
	}
	if v, ok := q.Peek(); ok {
		t.Fatalf("Peek on empty = (%d, true), want ok=false", v)
	}
}

func TestFIFOOrder(t *testing.T) {
	var q LinkedListQueue
	for i := 1; i <= 3; i++ {
		q.Enqueue(i)
		if q.Size() != i {
			t.Fatalf("after Enqueue(%d): Size=%d, want %d", i, q.Size(), i)
		}
	}
	if v, ok := q.Peek(); !ok || v != 1 {
		t.Fatalf("Peek = (%d, %v), want (1, true)", v, ok)
	}
	for want := 1; want <= 3; want++ {
		if v, ok := q.Dequeue(); !ok || v != want {
			t.Fatalf("Dequeue = (%d, %v), want (%d, true)", v, ok, want)
		}
		if q.Size() != 3-want {
			t.Fatalf("Size = %d, want %d", q.Size(), 3-want)
		}
	}
	if !q.IsEmpty() {
		t.Fatal("queue should be empty after dequeue everything")
	}
}

// Catches the classic bug of not resetting tail when the last node is removed.
func TestReuseAfterEmpty(t *testing.T) {
	var q LinkedListQueue
	q.Enqueue(1)
	q.Dequeue()
	q.Enqueue(2)
	q.Enqueue(3)
	for _, want := range []int{2, 3} {
		if v, ok := q.Dequeue(); !ok || v != want {
			t.Fatalf("Dequeue = (%d, %v), want (%d, true)", v, ok, want)
		}
	}
	if _, ok := q.Dequeue(); ok {
		t.Fatal("Dequeue should fail on empty queue")
	}
}

func TestInterleaved(t *testing.T) {
	var q LinkedListQueue
	var got []int
	q.Enqueue(1)
	q.Enqueue(2)
	v, _ := q.Dequeue()
	got = append(got, v)
	q.Enqueue(3)
	q.Enqueue(4)
	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		got = append(got, v)
	}
	want := []int{1, 2, 3, 4}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}
