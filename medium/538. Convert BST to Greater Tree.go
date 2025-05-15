package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + sumTree(root.Left) + sumTree(root.Right)
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := root.Val + sumTree(root.Left) + sumTree(root.Right)
	visit := map[*TreeNode]bool{}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		if curr.Left != nil && !visit[curr.Left] {
			stack = append(stack, curr.Left)
			continue
		}
		sum, curr.Val = sum-curr.Val, sum
		visit[curr] = true
		stack = stack[:len(stack)-1]
		if curr.Right != nil && !visit[curr.Right] {
			stack = append(stack, curr.Right)
		}
	}
	return root
}
