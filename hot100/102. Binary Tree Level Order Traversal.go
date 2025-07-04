package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var trace func(root *TreeNode, level int)
	ans := make([][]int, 0)
	trace = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(ans) <= level {
			ans = append(ans, []int{})
		}
		ans[level] = append(ans[level], root.Val)
		trace(root.Left, level+1)
		trace(root.Right, level+1)
	}
	trace(root, 0)
	return ans
}
