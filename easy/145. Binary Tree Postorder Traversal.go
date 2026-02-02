package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversalRecur260201(root *TreeNode) []int {
	ans := []int{}
	var backOrder func(node *TreeNode)
	backOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		backOrder(node.Left)
		backOrder(node.Right)
		ans = append(ans, node.Val)
	}

	backOrder(root)
	return ans
}

func postorderTraversalIter260201(root *TreeNode) []int {
	ans := []int{}
	stack1, stack2 := []*TreeNode{}, []*TreeNode{}
	curr := root
	for curr != nil {
		for curr != nil {
			stack1 = append(stack1, curr)
			curr = curr.Left
		}
		curr = stack1[len(stack1)-1]
		for curr.Right == nil || (len(stack2) > 0 && curr == stack2[len(stack2)-1]) {
			ans = append(ans, curr.Val)
			stack1 = stack1[:len(stack1)-1]
			if len(stack2) > 0 && curr == stack2[len(stack2)-1] {
				stack2 = stack2[:len(stack2)-1]
			}
			if len(stack1) == 0 {
				return ans
			}
			curr = stack1[len(stack1)-1]
		}
		if curr.Right != nil {
			stack2 = append(stack2, curr)
			curr = curr.Right
		}
	}
	return ans
}
