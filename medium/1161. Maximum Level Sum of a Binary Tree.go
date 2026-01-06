package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum260106(root *TreeNode) int {
	queue := []*TreeNode{root}
	ans := 0
	maxNum := -100000
	level := 0
	for len(queue) != 0 {
		level += 1
		curr := 0
		tmpQ := []*TreeNode{}
		for i := 0; i < len(queue); i++ {
			curr += queue[i].Val
			if queue[i].Left != nil {
				tmpQ = append(tmpQ, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmpQ = append(tmpQ, queue[i].Right)
			}
		}
		if curr > maxNum {
			ans = level
			maxNum = curr
		}
		queue = tmpQ
	}
	return ans
}
