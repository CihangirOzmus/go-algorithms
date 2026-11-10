package reverselinkedlist

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

func RecursivelyReverseList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	tail := RecursivelyReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tail
}
