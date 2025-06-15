package hot100
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    flag := (m+n)%2
    target1 := (m+n+1)/2
    ans1, ans2 := 0, 0
    for target1 != 0 {
        if m == 0 {
            ans1 = nums2[target1-1]
            nums2 = nums2[target1:]
            break
        }
        if n == 0 {
            ans1 = nums1[target1-1]
            nums1 = nums1[target1:]
            break
        }
        if target1 == 1 {
            if nums1[0] < nums2[0] {
                ans1 = nums1[0]
                nums1 = nums1[1:]
            } else {
                ans1 = nums2[0]
                nums2 = nums2[1:]
            }
            break
        }
        k1 := min(target1/2, m) - 1
        k2 := min(target1/2, n) - 1
        if nums1[k1] <= nums2[k2] {
            m -= k1 + 1
            target1 -= k1 + 1
            nums1 = nums1[k1+1:]
        } else {
            n -= k2 + 1
            target1 -= k2 + 1
            nums2 = nums2[k2+1:]
        }
    }
    if flag == 1 {
        return float64(ans1)
    }
    if len(nums1) == 0 {
        ans2 = nums2[0]
    } else if len(nums2) == 0 {
        ans2 = nums1[0]
    } else {
        ans2 = min(nums1[0], nums2[0])
    }
    return float64(ans1+ans2)/2
}