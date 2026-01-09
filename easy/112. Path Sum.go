package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum260109(root *TreeNode, targetSum int) bool {
	ans := false
	currSum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if ans || node == nil {
			return
		}
		currSum += node.Val
		if node.Left == nil && node.Right == nil && currSum == targetSum {
			ans = true
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		currSum -= node.Val
		return
	}

	dfs(root)
	return ans
}
