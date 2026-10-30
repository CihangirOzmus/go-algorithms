package linkedlist

type OptimalSinglyLinkedList struct {
	head *ListNode
	size int
}

func NewOptimalSinglyLinkedList() OptimalSinglyLinkedList {
	return OptimalSinglyLinkedList{head: &ListNode{val: 0}, size: 0}
}

func (s *OptimalSinglyLinkedList) getPrev(index int) *ListNode {
	cur := s.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (s *OptimalSinglyLinkedList) Get(index int) int {
	if index >= s.size {
		return -1
	}
	return s.getPrev(index).next.val
}

func (s *OptimalSinglyLinkedList) AddAtHead(val int) {
	s.AddAtIndex(0, val)
}

func (s *OptimalSinglyLinkedList) AddAtTail(val int) {
	s.AddAtIndex(s.size, val)
}

func (s *OptimalSinglyLinkedList) AddAtIndex(index int, val int) {
	if index > s.size {
		return
	}
	prev := s.getPrev(index)
	node := &ListNode{val: val, next: prev.next}
	prev.next = node
	s.size++
}

func (s *OptimalSinglyLinkedList) DeleteAtIndex(index int) {
	if index >= s.size {
		return
	}
	prev := s.getPrev(index)
	prev.next = prev.next.next
	s.size--
}
