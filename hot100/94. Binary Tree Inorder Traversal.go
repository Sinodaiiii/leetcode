package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	visited := map[*TreeNode]bool{}
	ans := make([]int, 0)
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		if curr.Left != nil && !visited[curr.Left] {
			stack = append(stack, curr.Left)
			continue
		}
		ans = append(ans, curr.Val)
		visited[curr] = true
		stack = stack[:len(stack)-1]
		if curr.Right != nil && !visited[curr.Right] {
			stack = append(stack, curr.Right)
		}
	}
	return ans
}
