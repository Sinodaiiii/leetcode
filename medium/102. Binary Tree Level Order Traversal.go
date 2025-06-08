package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	lStack := []int{1}
	ans := make([][]int, 0)
	for len(stack) != 0 {
		curr := stack[0]
		level := lStack[0]
		stack, lStack = stack[1:], lStack[1:]
		if len(ans) < level {
			ans = append(ans, []int{})
		}
		ans[level-1] = append(ans[level-1], curr.Val)
		if curr.Left != nil {
			stack = append(stack, curr.Left)
			lStack = append(lStack, level+1)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
			lStack = append(lStack, level+1)
		}
	}
	return ans
}
