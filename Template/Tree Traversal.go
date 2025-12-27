package Template

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	ans := []int{}
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		ans = append(ans, stack[len(stack)-1].Val)
		curr = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ans
}
