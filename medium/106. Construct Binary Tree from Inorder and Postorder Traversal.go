package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	curr := &TreeNode{postorder[n-1], nil, nil}
	pivot := -1
	for i := 0; i < n; i++ {
		if inorder[i] == postorder[n-1] {
			pivot = i
			break
		}
	}
	curr.Left = buildTree(inorder[:pivot], postorder[:pivot])
	curr.Right = buildTree(inorder[pivot+1:], postorder[pivot:])
	return curr
}
