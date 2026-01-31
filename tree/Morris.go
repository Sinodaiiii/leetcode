package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func morrisInorder(root *TreeNode) []int {
	var res []int
	curr := root

	for curr != nil {
		if curr.Left == nil {
			// 1. 没有左孩子，直接收录，往右走
			res = append(res, curr.Val)
			curr = curr.Right
		} else {
			// 2. 有左孩子，先找前驱节点 (predecessor)
			mostRight := curr.Left
			// 注意：寻找终点的条件增加了 mostRight.Right != curr
			//这是为了区分是“第一次来搭桥”还是“第二次来拆桥”
			for mostRight.Right != nil && mostRight.Right != curr {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				// 情况 A：第一次来到这里，搭桥
				mostRight.Right = curr
				curr = curr.Left
			} else {
				// 情况 B：桥已经搭好了，说明左子树遍历完毕，拆桥
				mostRight.Right = nil
				res = append(res, curr.Val) // 中序遍历在这里收录节点
				curr = curr.Right
			}
		}
	}
	return res
}

func morrisPreorder(root *TreeNode) []int {
	var res []int
	curr := root

	for curr != nil {
		if curr.Left == nil {
			// 1. 没有左孩子，直接收录，往右走
			res = append(res, curr.Val)
			curr = curr.Right
		} else {
			// 2. 有左孩子，先找前驱节点 (predecessor)
			mostRight := curr.Left
			// 注意：寻找终点的条件增加了 mostRight.Right != curr
			//这是为了区分是“第一次来搭桥”还是“第二次来拆桥”
			for mostRight.Right != nil && mostRight.Right != curr {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				// 情况 A：第一次来到这里，搭桥
				mostRight.Right = curr
				res = append(res, curr.Val) // <--- 前序在这里收录！先收录，再往左走
				curr = curr.Left
			} else {
				// 情况 B：第二次来到这里，拆桥
				mostRight.Right = nil
				// res = append(res, curr.Val) // <--- 中序在这里收录
				curr = curr.Right
			}
		}
	}
	return res
}
