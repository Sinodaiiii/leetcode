package medium

import "math"

type Node struct {
	value  int
	length int
	left   *Node
	right  *Node
}

func lengthOfLIS(nums []int) int {
	root := &Node{nums[0], 1, nil, nil}
	for i, value := range nums {
		if i == 0 {
			continue
		}
		converge(root, value)
	}
	return dfsMAX(root, math.MaxInt32)
}

func converge(root *Node, target int) {
	var same, father *Node = nil, nil
	curr := root
	for curr != nil {
		if target == curr.value {
			same = curr
			curr = curr.left
		} else if target > curr.value {
			father = curr
			curr = curr.right
		} else if target < curr.value {
			father = curr
			curr = curr.left
		}
	}
	if same != nil {
		same.length = max(same.length, dfsMAX(root, same.value)+1)
		return
	}
	curr = &Node{target, dfsMAX(root, target) + 1, nil, nil}
	if curr.value < father.value {
		father.left = curr
	} else {
		father.right = curr
	}
}

func dfsMAX(root *Node, maxv int) int {
	if root == nil {
		return 0
	}
	if root.value >= maxv {
		return max(dfsMAX(root.left, maxv), 0)
	}
	return max(dfsMAX(root.left, maxv), dfsMAX(root.right, maxv), root.length)
}
