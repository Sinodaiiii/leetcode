package medium

func subsetsWithDup(nums []int) [][]int {
    num := make([]int, 0)
    dict := map[int]int{}
    for i:=0; i<len(nums); i++ {
        if dict[nums[i]] == 0 {
            num = append(num, nums[i])
        }
        dict[nums[i]] += 1
    }
    n := len(num)
    ans := make([][]int, 0)
    cur := make([]int, 0)
    var dfs func(level int)
    dfs = func(level int){
        if level == n {
            tmp := make([]int, len(cur))
            copy(tmp, cur)
            ans = append(ans, tmp)
            return
        }
        dfs(level+1)
        for i:=1; i<=dict[num[level]]; i++ {
            cur = append(cur, num[level])
            dfs(level+1)
        }
        cur = cur[:len(cur)-dict[num[level]]]
    }
    dfs(0)
    return ans
}