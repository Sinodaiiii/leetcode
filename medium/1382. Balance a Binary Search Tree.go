package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST260209(root *TreeNode) *TreeNode {
	tree := []*TreeNode{}
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		tree = append(tree, node)
		inOrder(node.Right)
	}
	inOrder(root)

	var getTree func(l, r int) *TreeNode
	getTree = func(l, r int) *TreeNode {
		if l == r {
			tree[l].Left, tree[l].Right = nil, nil
			return tree[l]
		} else if l > r {
			return nil
		}
		mid := (l + r) / 2
		tree[mid].Left = getTree(l, mid-1)
		tree[mid].Right = getTree(mid+1, r)
		return tree[mid]
	}

	return getTree(0, len(tree)-1)
}
