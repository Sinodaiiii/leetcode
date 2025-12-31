package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	ans := 0
	curr := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		curr = curr*10 + root.Val
		if root.Left == nil && root.Right == nil {
			ans += curr
		} else {
			if root.Left != nil {
				dfs(root.Left)
			}
			if root.Right != nil {
				dfs(root.Right)
			}
		}
		curr = curr / 10
	}

	dfs(root)
	return ans
}

func sumNumbers251231(root *TreeNode) int {
	ans := 0
	curr := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		curr = curr*10 + node.Val
		if node.Left == nil && node.Right == nil {
			ans += curr
		} else {
			if node.Left != nil {
				dfs(node.Left)
			}
			if node.Right != nil {
				dfs(node.Right)
			}
		}
		curr = curr / 10
	}

	dfs(root)
	return ans
}
