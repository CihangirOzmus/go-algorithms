package linkedlist

type doublyNode struct {
	val        int
	prev, next *doublyNode
}

// Doubly is a doubly linked list. It uses sentinel head and tail nodes, so
// insert/delete never need nil checks, and walks from whichever end is closer
// to the index.
type Doubly struct {
	head, tail *doublyNode
	size       int
}

func NewDoubly() *Doubly {
	head, tail := &doublyNode{}, &doublyNode{}
	head.next, tail.prev = tail, head
	return &Doubly{head: head, tail: tail}
}

// nodeAt returns the node at index, where index == size means the tail sentinel.
func (s *Doubly) nodeAt(index int) *doublyNode {
	if index < s.size/2 {
		cur := s.head.next
		for range index {
			cur = cur.next
		}
		return cur
	}
	cur := s.tail
	for range s.size - index {
		cur = cur.prev
	}
	return cur
}

func (s *Doubly) Get(index int) int {
	if index < 0 || index >= s.size {
		return -1
	}
	return s.nodeAt(index).val
}

func (s *Doubly) AddAtHead(val int) {
	s.AddAtIndex(0, val)
}

func (s *Doubly) AddAtTail(val int) {
	s.AddAtIndex(s.size, val)
}

func (s *Doubly) AddAtIndex(index int, val int) {
	if index < 0 || index > s.size {
		return
	}
	next := s.nodeAt(index)
	node := &doublyNode{val: val, prev: next.prev, next: next}
	next.prev.next = node
	next.prev = node
	s.size++
}

func (s *Doubly) DeleteAtIndex(index int) {
	if index < 0 || index >= s.size {
		return
	}
	node := s.nodeAt(index)
	node.prev.next = node.next
	node.next.prev = node.prev
	s.size--
}
