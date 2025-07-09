package hot100

func findDuplicate(nums []int) int {
	bit := [2000]int{}
	for i := 0; i < len(nums); i++ {
		bi, bj := nums[i]/60, nums[i]%60
		if bit[bi]&(1<<bj) != 0 {
			return nums[i]
		}
		bit[bi] += 1 << bj
	}
	return 0
}
