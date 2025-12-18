package medium

func widthOfBinaryTree(root *TreeNode) int {
	stack := []*TreeNode{root}
	index := []int{1}
	ans := 0
	for len(stack) != 0 {
		ans = max(ans, index[len(index)-1]-index[0]+1)
		tmpStack := []*TreeNode{}
		tmpIndex := []int{}
		for i, curr := range stack {
			if curr.Left != nil {
				tmpStack = append(tmpStack, curr.Left)
				tmpIndex = append(tmpIndex, index[i]*2)
			}
			if curr.Right != nil {
				tmpStack = append(tmpStack, curr.Right)
				tmpIndex = append(tmpIndex, index[i]*2+1)
			}
		}
		stack, index = tmpStack, tmpIndex
	}
	return ans
}
