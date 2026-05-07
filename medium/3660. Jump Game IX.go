package medium

func maxValue260507(nums []int) []int {
	n := len(nums)
	leftMax := nums[0]
	leftG := make([]int, n)
	// for i := range leftG {
	//     leftG[i] = i
	// }
	stackG := make([]int, 0, n)
	findStack := func(num int) int {
		l, r := 0, len(stackG)-1
		for l < r {
			mid := (l + r) / 2
			if nums[stackG[mid]] <= num {
				l = mid + 1
			} else {
				if l == mid {
					break
				}
				r = mid
			}
		}
		return stackG[l]
	}
	ans := make([]int, len(nums))
	for i, num := range nums {
		if num >= leftMax {
			leftMax = num
		}
		ans[i] = leftMax
		if len(stackG) == 0 || nums[stackG[len(stackG)-1]] < num {
			stackG = append(stackG, i)
		}
		leftG[i] = findStack(num)
	}
	// fmt.Println(ans)
	// fmt.Println(leftG)
	queue := make([][2]int, 0, n)
	for i := len(nums) - 1; i >= 0; i-- {
		for len(queue) != 0 && queue[0][0] > i {
			queue = queue[1:]
		}
		if len(queue) > 0 {
			ans[i] = max(ans[i], queue[0][1])
		}
		if len(queue) == 0 || leftG[i] < queue[len(queue)-1][0] {
			queue = append(queue, [2]int{leftG[i], ans[i]})
		}
	}
	return ans
}
