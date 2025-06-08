package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	flag := 1
	ans := [][]int{}
	for len(stack) != 0 {
		tmpStack := []*TreeNode{}
		ans = append(ans, []int{})
		for i := len(stack) - 1; i >= 0; i-- {
			ans[len(ans)-1] = append(ans[len(ans)-1], stack[i].Val)
			var first, next *TreeNode
			if flag > 0 {
				first, next = stack[i].Left, stack[i].Right
			} else {
				first, next = stack[i].Right, stack[i].Left
			}
			if first != nil {
				tmpStack = append(tmpStack, first)
			}
			if next != nil {
				tmpStack = append(tmpStack, next)
			}
		}
		stack = tmpStack
		flag *= -1
	}
	return ans
}
