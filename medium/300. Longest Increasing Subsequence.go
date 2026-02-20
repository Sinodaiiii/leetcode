package medium

import "math"

type Node300 struct {
	value  int
	length int
	left   *Node300
	right  *Node300
}

func lengthOfLIS(nums []int) int {
	root := &Node300{nums[0], 1, nil, nil}
	for i, value := range nums {
		if i == 0 {
			continue
		}
		converge(root, value)
	}
	return dfsMAX(root, math.MaxInt32)
}

func converge(root *Node300, target int) {
	var same, father *Node300 = nil, nil
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
	curr = &Node300{target, dfsMAX(root, target) + 1, nil, nil}
	if curr.value < father.value {
		father.left = curr
	} else {
		father.right = curr
	}
}

func dfsMAX(root *Node300, maxv int) int {
	if root == nil {
		return 0
	}
	if root.value >= maxv {
		return max(dfsMAX(root.left, maxv), 0)
	}
	return max(dfsMAX(root.left, maxv), dfsMAX(root.right, maxv), root.length)
}

func lengthOfLIS260220(nums []int) int {
	ans := make([]int, 0, len(nums))

	for i, num := range nums {
		target := -1
		l, r := 0, len(ans)-1
		for l <= r {
			mid := (l + r) / 2
			if nums[ans[mid]] == num {
				target = mid
				break
			} else if nums[ans[mid]] < num {
				target = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if target != -1 && nums[ans[target]] == num {
			continue
		} else {
			target += 1
			if target >= len(ans) {
				ans = append(ans, i)
			} else {
				if nums[ans[target]] > num {
					ans[target] = i
				}
			}
		}
	}

	return len(ans)
}
