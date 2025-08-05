package medium

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	dict12 := map[int]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dict12[nums1[i]+nums2[j]] += 1
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += dict12[-(nums3[i] + nums4[j])]
		}
	}
	return ans
}
