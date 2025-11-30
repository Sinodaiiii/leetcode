package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode = nil
	var findChild func(node *TreeNode) int
	findChild = func(node *TreeNode) int {
		if node == nil || ans != nil {
			return 0
		}
		l := findChild(node.Left)
		r := findChild(node.Right)
		curr := 0
		if node == p || node == q {
			curr = 1
		}
		if l+r+curr == 2 {
			ans = node
			return 3
		}
		return l + r + curr
	}
	findChild(root)
	return ans
}
