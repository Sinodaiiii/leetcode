package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root.Left == nil || root.Right == nil {
		return 1
	}
	ans := 0
	visited := map[*TreeNode]bool{}
	leftStack, rightStack := []*TreeNode{root.Left}, []*TreeNode{root.Right}
	li, ri := []int{1, 2}, []int{1, 3}
	depth := 1
	i := 2
	for len(leftStack) != 0 {
		curr := leftStack[len(leftStack)-1]
		if curr.Left != nil && visited[curr.Left] != true {
			leftStack = append(leftStack, curr.Left)
			depth += 1
			i = i * 2
			visited[curr.Left] = true
			if len(li) > depth {
				li[depth] = min(li[depth], i)
			} else {
				li = append(li, i)
			}
		} else if curr.Right != nil && visited[curr.Right] != true {
			leftStack = append(leftStack, curr.Right)
			depth += 1
			i = i*2 + 1
			visited[curr.Right] = true
			if len(li) > depth {
				li[depth] = min(li[depth], i)
			} else {
				li = append(li, i)
			}
		} else {
			leftStack = leftStack[0 : len(leftStack)-1]
			depth -= 1
			i = i / 2
		}
	}
	depth = 1
	i = 3
	for len(rightStack) != 0 {
		curr := rightStack[len(rightStack)-1]
		if curr.Left != nil && visited[curr.Left] != true {
			rightStack = append(rightStack, curr.Left)
			depth += 1
			i = i * 2
			visited[curr.Left] = true
			if len(ri) > depth {
				ri[depth] = max(ri[depth], i)
			} else {
				ri = append(ri, i)
			}
		} else if curr.Right != nil && visited[curr.Right] != true {
			rightStack = append(rightStack, curr.Right)
			depth += 1
			i = i*2 + 1
			visited[curr.Right] = true
			if len(ri) > depth {
				ri[depth] = max(ri[depth], i)
			} else {
				ri = append(ri, i)
			}
		} else {
			rightStack = rightStack[0 : len(rightStack)-1]
			depth -= 1
			i = i / 2
		}
	}
	for i = 0; i < min(len(li), len(ri)); i++ {
		ans = max(ans, ri[i]-li[i]+1)
	}
	ansl := 0
	if root.Left != nil {
		ansl = widthOfBinaryTree(root.Left)
	}
	ansr := 0
	if root.Right != nil {
		ansr = widthOfBinaryTree(root.Right)
	}
	return max(ans, ansl, ansr)
}
