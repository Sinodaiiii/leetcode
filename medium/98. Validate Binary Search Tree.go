package medium

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var check func(node *TreeNode) (bool, int, int)
	check = func(node *TreeNode) (bool, int, int) {
		if node == nil {
			return true, math.MaxInt32 * 2, math.MinInt32 * 2
		}
		la, lm, ll := check(node.Left)
		ra, rm, rl := check(node.Right)
		return la && ra && ll < node.Val && rm > node.Val, min(lm, node.Val), max(node.Val, rl)
	}

	ret, _, _ := check(root)
	return ret
}
