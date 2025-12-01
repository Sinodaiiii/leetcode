package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	ans := -1
	var sonNum func(fatherNum int, node *TreeNode) int
	sonNum = func(fatherNum int, node *TreeNode) int {
		if ans >= 0 || node == nil {
			return 0
		}
		l := sonNum(fatherNum, node.Left)
		if ans < 0 && l+fatherNum == k-1 {
			ans = node.Val
		}
		r := sonNum(l+fatherNum+1, node.Right)
		return l + r + 1
	}
	sonNum(0, root)
	return ans
}
