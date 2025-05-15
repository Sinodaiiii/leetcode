package medium

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	roby, robn := robSubTree(root)
	return max(roby, robn)
}

func robSubTree(root *TreeNode) (roby, robn int) {
	if root == nil {
		return math.MinInt32, 0
	}
	lroby, lrobn := robSubTree(root.Left)
	rroby, rrobn := robSubTree(root.Right)
	roby = root.Val + lrobn + rrobn
	robn = max(lroby+rroby, lroby+rrobn, lrobn+rroby, lrobn+rrobn)
	return roby, robn
}
