package medium

func flatten(root *TreeNode) {
	flattenTree(root)
}

func flattenTree(root *TreeNode) (tail *TreeNode) {
	if root == nil {
		return nil
	}
	ltail := flattenTree(root.Left)
	rtail := flattenTree(root.Right)
	l, r, tail := root.Left, root.Right, root
	if ltail != nil {
		tail.Right, tail = l, ltail
	}
	if rtail != nil {
		tail.Right, tail = r, rtail
	}
	root.Left = nil
	return
}
