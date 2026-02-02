package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree260202(root *TreeNode, subRoot *TreeNode) bool {
	var check func(node1, node2 *TreeNode) bool
	check = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if (node1 == nil && node2 != nil) || (node2 == nil && node1 != nil) || (node1.Val != node2.Val) {
			return false
		}
		return check(node1.Left, node2.Left) && check(node1.Right, node2.Right)
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
		if curr.Val == subRoot.Val {
			if check(curr, subRoot) {
				return true
			}
		}
	}
	return false
}
