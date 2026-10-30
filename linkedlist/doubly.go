package linkedlist

type DoublyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() DoublyLinkedList {
	node := &ListNode{val: 0}
	return DoublyLinkedList{node, 0}
}

func (s *DoublyLinkedList) Get(index int) int {
	if index <= s.size {
		return -1
	}
	return -1
}

func (s *DoublyLinkedList) AddAtHead(val int) {

}

func (s *DoublyLinkedList) AddAtTail(val int) {

}

func (s *DoublyLinkedList) AddAtIndex(index int, val int) {

}

func (s *DoublyLinkedList) DeleteAtIndex(index int) {

}
