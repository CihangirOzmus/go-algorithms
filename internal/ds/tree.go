package ds

// Null marks a missing node in level-order input, like "null" on LeetCode.
const Null = -1 << 63

// TreeNode is LeetCode's binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTree builds a binary tree from LeetCode-style level-order values,
// using Null for missing nodes: NewTree(1, Null, 2, 3).
func NewTree(vals ...int) *TreeNode {
	if len(vals) == 0 || vals[0] == Null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(vals); {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != Null {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != Null {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// TreeToSlice returns the level-order values of the tree, using Null for
// missing nodes and trimming trailing Nulls (same format as LeetCode).
func TreeToSlice(root *TreeNode) []int {
	var out []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			out = append(out, Null)
			continue
		}
		out = append(out, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(out) > 0 && out[len(out)-1] == Null {
		out = out[:len(out)-1]
	}
	return out
}
