package main

import "github.com/cihangirozmus/go-algorithms/linkedlists"

func main() {
	node := linkedlists.NewListNode(-1)
	s := linkedlists.NewSinglyLinkedList()
	s.Head, s.Tail = node, node

	s.InsertEnd(1)
	s.InsertEnd(2)
	s.InsertEnd(3)
	s.InsertEnd(4)
	s.InsertEnd(5)
	s.InsertEnd(6)
	s.Print()

	s.Remove(3)
	s.Print()
}
