package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum113260109(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	curr := []int{}
	currSum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		curr = append(curr, node.Val)
		currSum += node.Val
		if node.Left == nil && node.Right == nil && currSum == targetSum {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			ans = append(ans, tmp)
			currSum -= node.Val
			curr = curr[:len(curr)-1]
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		curr = curr[:len(curr)-1]
		currSum -= node.Val
		return
	}

	dfs(root)
	return ans
}
