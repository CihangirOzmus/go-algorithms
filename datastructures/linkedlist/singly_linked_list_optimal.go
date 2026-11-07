package linkedlist

type SinglyOptimal struct {
	head *singlyNode
	size int
}

func NewSinglyOptimal() *SinglyOptimal {
	return &SinglyOptimal{head: &singlyNode{}}
}

func (s *SinglyOptimal) getPrev(index int) *singlyNode {
	cur := s.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (s *SinglyOptimal) Get(index int) int {
	if index >= s.size {
		return -1
	}
	return s.getPrev(index).next.val
}

func (s *SinglyOptimal) AddAtHead(val int) {
	s.AddAtIndex(0, val)
}

func (s *SinglyOptimal) AddAtTail(val int) {
	s.AddAtIndex(s.size, val)
}

func (s *SinglyOptimal) AddAtIndex(index int, val int) {
	if index > s.size {
		return
	}
	prev := s.getPrev(index)
	node := &singlyNode{val: val, next: prev.next}
	prev.next = node
	s.size++
}

func (s *SinglyOptimal) DeleteAtIndex(index int) {
	if index >= s.size {
		return
	}
	prev := s.getPrev(index)
	prev.next = prev.next.next
	s.size--
}
