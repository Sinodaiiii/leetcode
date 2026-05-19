package easy

func getCommon260519(nums1 []int, nums2 []int) int {
	p1, p2 := 0, 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			return nums1[p1]
		} else if nums1[p1] > nums2[p2] {
			p2 += 1
		} else {
			p1 += 1
		}
	}
	return -1
}
