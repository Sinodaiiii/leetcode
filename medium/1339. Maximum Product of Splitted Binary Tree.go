package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct260107(root *TreeNode) int {
	// ans := 0
	sum := make([]int, 0)
	total := 0
	var getSum func(node *TreeNode) int
	getSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		total += node.Val
		l := getSum(node.Left)
		r := getSum(node.Right)
		sum = append(sum, l+r+node.Val)
		return l + r + node.Val
	}

	getSum(root)
	maxV := int64(0)
	for _, v := range sum {
		maxV = max(maxV, int64(v)*int64(total-v))
	}
	return int(maxV % 1000000007)
}
