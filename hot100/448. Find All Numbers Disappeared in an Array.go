package hot100

func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    for i:=0; i<n; i++ {
        if nums[i] != 0 {
            j := nums[i]
            for j!=0 {
                j, nums[j-1] = nums[j-1], 0
            }
        }
    }
    ans := []int{}
    for i:=0; i<n; i++ {
        if nums[i] != 0 {
            ans = append(ans, i+1)
        }
    }
    return ans
}
