package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	curr := preorder[0]
	root := &TreeNode{curr, nil, nil}
	inIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == curr {
			inIndex = i
			break
		}
	}
	if inIndex > 0 {
		root.Left = buildTree(preorder[1:inIndex+1], inorder[:inIndex])
	}
	if inIndex < len(inorder)-1 {
		root.Right = buildTree(preorder[inIndex+1:], inorder[inIndex+1:])
	}
	return root
}
