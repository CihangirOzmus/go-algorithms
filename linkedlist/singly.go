package linkedlist

type SinglyLinkedList struct {
	head *ListNode
	size int
}

// NewSinglyLinkedList is the constructor
func NewSinglyLinkedList() SinglyLinkedList {
	node := &ListNode{val: 0}
	return SinglyLinkedList{node, 0}
}

func (s *SinglyLinkedList) Get(index int) int {
	if index >= s.size {
		return -1
	}
	cur := s.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (s *SinglyLinkedList) AddAtHead(val int) {
	node := &ListNode{val: val}
	node.next = s.head.next
	s.head.next = node
	s.size++
}

func (s *SinglyLinkedList) AddAtTail(val int) {
	node := &ListNode{val: val}
	cur := s.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
	s.size++
}

func (s *SinglyLinkedList) AddAtIndex(index int, val int) {
	if index > s.size {
		return
	}
	cur := s.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	node := &ListNode{val: val}
	node.next = cur.next
	cur.next = node
	s.size++
}

func (s *SinglyLinkedList) DeleteAtIndex(index int) {
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
