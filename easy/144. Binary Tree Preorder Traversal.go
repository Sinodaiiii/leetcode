package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal260105(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	ans := []int{}
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, curr.Val)
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return ans
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{root.Val}
	ret = append(ret, preorderTraversal(root.Left)...)
	ret = append(ret, preorderTraversal(root.Right)...)
	return ret
}
