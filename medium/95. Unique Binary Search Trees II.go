package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	var generate func(l, r int) []*TreeNode

	generate = func(l, r int) []*TreeNode {
		if r < l {
			return []*TreeNode{nil}
		} else if r == l {
			return []*TreeNode{&TreeNode{l, nil, nil}}
		}

		ret := []*TreeNode{}
		for i := l; i <= r; i++ {
			lTree := generate(l, i-1)
			rTree := generate(i+1, r)
			for j := 0; j < len(lTree); j++ {
				for k := 0; k < len(rTree); k++ {
					head := &TreeNode{i, lTree[j], rTree[k]}
					ret = append(ret, head)
				}
			}
		}
		return ret
	}

	return generate(1, n)
}
