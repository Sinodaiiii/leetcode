package medium

func maxDistance260419(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	ans := 0
	for i < len(nums1) && j < len(nums2) {
		for j < len(nums2) && nums2[j] >= nums1[i] {
			ans = max(ans, j-i)
			j += 1
		}
		i += 1
	}
	return ans
}
