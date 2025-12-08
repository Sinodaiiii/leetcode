package medium

type TreeNode538 struct {
	Val   int
	Left  *TreeNode538
	Right *TreeNode538
}

func sumTree(root *TreeNode538) int {
	if root == nil {
		return 0
	}
	return root.Val + sumTree(root.Left) + sumTree(root.Right)
}

func convertBST(root *TreeNode538) *TreeNode538 {
	if root == nil {
		return root
	}
	sum := root.Val + sumTree(root.Left) + sumTree(root.Right)
	visit := map[*TreeNode538]bool{}
	stack := []*TreeNode538{root}
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

func convertBST2(root *TreeNode538) *TreeNode538 {
	if root == nil {
		return nil
	}
	var checkSon func(node *TreeNode538, fatherVal int) int
	checkSon = func(node *TreeNode538, fatherVal int) int {
		if node.Right != nil {
			node.Val += checkSon(node.Right, fatherVal)
		} else {
			node.Val += fatherVal
		}
		if node.Left != nil {
			return checkSon(node.Left, node.Val)
		} else {
			return node.Val
		}
	}

	checkSon(root, 0)
	return root
}
