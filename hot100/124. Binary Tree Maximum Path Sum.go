package hot100

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var getMax func(root *TreeNode) int
	getMax = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lMax := getMax(root.Left)
		rMax := getMax(root.Right)
		ret := max(root.Val, root.Val+lMax, root.Val+rMax)
		curr := root.Val + lMax + rMax
		ans = max(curr, ret, ans)
		return ret
	}

	getMax(root)
	return ans

}
