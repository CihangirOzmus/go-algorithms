package linkedlist

type Singly struct {
	head *singlyNode
	size int
}

// NewSingly returns an empty singly linked list with a sentinel head.
func NewSingly() *Singly {
	return &Singly{head: &singlyNode{}}
}

func (s *Singly) Get(index int) int {
	if index >= s.size {
		return -1
	}
	cur := s.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (s *Singly) AddAtHead(val int) {
	node := &singlyNode{val: val}
	node.next = s.head.next
	s.head.next = node
	s.size++
}

func (s *Singly) AddAtTail(val int) {
	node := &singlyNode{val: val}
	cur := s.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
	s.size++
}

func (s *Singly) AddAtIndex(index int, val int) {
	if index > s.size {
		return
	}
	cur := s.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	node := &singlyNode{val: val}
	node.next = cur.next
	cur.next = node
	s.size++
}

func (s *Singly) DeleteAtIndex(index int) {
	if index >= s.size {
		return
	}
	cur := s.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	s.size--
}
