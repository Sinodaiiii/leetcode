package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	queue := []*TreeNode{root}
	length := 1
	for length != 0 {
		nextLength := 0
		currAns := []int{}
		for length > 0 {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
				nextLength += 1
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
				nextLength += 1
			}
			currAns = append(currAns, queue[0].Val)
			queue = queue[1:]
			length -= 1
		}
		tmp := [][]int{currAns}
		ans = append(tmp, ans...)
		length = nextLength
	}
	return ans
}
