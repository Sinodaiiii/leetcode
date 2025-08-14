package medium

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	length := make([][2]int, n)
	for i := 0; i < n; i++ {
		length[i] = [2]int{1, 1}
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if length[j][0] >= length[i][0] {
					length[i] = [2]int{length[j][0] + 1, length[j][1]}
				} else if length[j][0] == length[i][0]-1 {
					length[i][1] += length[j][1]
				}
			}
		}
	}
	ans := 0
	maxL := 1
	for i := 0; i < n; i++ {
		if length[i][0] == maxL {
			ans += length[i][1]
		} else if length[i][0] > maxL {
			maxL = length[i][0]
			ans = length[i][1]
		}
	}
	return ans
}
