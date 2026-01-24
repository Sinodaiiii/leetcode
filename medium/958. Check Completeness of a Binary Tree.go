package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree260125(root *TreeNode) bool {
	var checkTree func(node *TreeNode) (int, bool)
	checkTree = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		ld, lb := checkTree(node.Left)
		rd, rb := checkTree(node.Right)
		// fmt.Println(node.Val, ld, lb, rd, rb)
		if ld < 0 || rd < 0 || ld < rd || ld > rd+1 {
			return -1, false
		}
		if ld == rd {
			if lb == true {
				return ld + 1, rb == true
			}
			return -1, false
		} else {
			if rb == true {
				return ld + 1, false
			}
			return -1, false
		}
	}

	depth, _ := checkTree(root)
	return depth > 0
}
