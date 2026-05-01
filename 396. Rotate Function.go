package main

func maxRotateFunction260501(nums []int) int {
	ans, total := 0, 0
	n := len(nums)
	for i, num := range nums {
		ans += i * num
		total += num
	}
	// fmt.Println(ans, total)
	curr := ans
	for i := n - 1; i >= 0; i-- {
		curr = curr - n*nums[i] + total
		ans = max(ans, curr)
		// fmt.Println(curr, ans)
	}
	return ans
}
