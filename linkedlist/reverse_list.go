package linkedlist

type Node struct {
	Val  int
	Next *Node
}

func ReverseList(head *Node) *Node {
	var prev *Node
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
