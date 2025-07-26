package easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	j1, j2 := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if j1 == -1 || j2 == -1 {
			if j1 == -1 {
				for j2 >= 0 {
					nums1[j2] = nums2[j2]
					j2 -= 1
				}
			}
			return
		}
		if nums1[j1] >= nums2[j2] {
			nums1[i] = nums1[j1]
			j1 -= 1
		} else {
			nums1[i] = nums2[j2]
			j2 -= 1
		}
	}
}
