package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getNum260205(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + getNum260205(node.Left) + getNum260205(node.Right)
}

func findTargetNode260205(root *TreeNode, cnt int) int {
	r := getNum260205(root.Right)
	if r == cnt-1 {
		return root.Val
	} else if r >= cnt {
		return findTargetNode260205(root.Right, cnt)
	} else {
		return findTargetNode260205(root.Left, cnt-r-1)
	}
}
