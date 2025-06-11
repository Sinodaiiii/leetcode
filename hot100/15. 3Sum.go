package hot100

import "sort"

func threeSum(nums []int) [][]int {
    ans := [][]int{}
    sort.Slice(nums, func(i, j int)bool {return nums[i]<nums[j]})
    n := len(nums)
    for i:=0; i<n; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        l, r, target := i+1, n-1, -nums[i]
        for l < r {
            if nums[l] + nums[r] > target {
                r -= 1
            } else if nums[l] + nums[r] < target {
                l += 1
            } else {
                ans = append(ans, []int{nums[i], nums[l], nums[r]})
                l += 1
                for l < r && nums[l] == nums[l-1] {
                    l += 1
                }
			}
        }
    }
    return ans
}
