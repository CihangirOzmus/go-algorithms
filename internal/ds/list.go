// Package ds holds the shared data structures LeetCode predefines
// (ListNode, TreeNode) plus helpers for building them in tests.
package ds

// ListNode is LeetCode's singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewList builds a linked list from vals and returns its head (nil if empty).
func NewList(vals ...int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// ListToSlice returns the values of the list starting at head.
func ListToSlice(head *ListNode) []int {
	var out []int
	for ; head != nil; head = head.Next {
		out = append(out, head.Val)
	}
	return out
}
