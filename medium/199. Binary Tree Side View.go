package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, deep int, ret []int) []int {
	if root == nil {
		return ret
	}
	deep += 1
	if len(ret) < deep {
		ret = append(ret, root.Val)
	}
	ret = dfs(root.Right, deep, ret)
	ret = dfs(root.Left, deep, ret)
	deep -= 1
	return ret
}

func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	deep := 0
	ret = dfs(root, deep, ret)
	return ret
}
