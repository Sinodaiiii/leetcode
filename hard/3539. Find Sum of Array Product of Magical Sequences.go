package hard

import (
	"fmt"
	"math"
)

func magicalSum(m int, k int, nums []int) int {
	var tt, mod int64
	tt = int64(1)
	for i := 0; i < len(nums); i++ {
		tt *= int64(nums[i])
	}
	mod = int64(math.Pow(10, 9) + 7)
	fmt.Println(tt % mod)
	return 0
	// dict := map[int]int{}
	// keys := []int{}
	// for i:=0; i<len(nums); i++ {
	//     if dict[nums[i]] == 0 {
	//         keys = append(keys, nums[i])
	//     }
	//     dict[nums[i]] += 1
	// }
	// curr := []int{}
	// ans := 0
	// n := len(keys)
	// left := k
	// var dfs func(index int)
	// fmt.Println(keys)
	// fmt.Println(dict)
	// dfs = func(index int) {
	//     fmt.Println(curr)
	//     if index >= n && len(curr) != m {
	//         return
	//     }
	//     if len(curr) == m {
	//         if left > 0 {
	//             return
	//         }
	//         tmp := 1
	//         for i:=0; i<m; i++ {
	//             tmp *= keys[curr[i]]
	//         }
	//         ans += tmp % mod
	//         return
	//     }
	//     left -= 1
	//     pivot := len(curr)
	//     for i:=0; i<dict[keys[index]] && (i+len(curr))<m; i++ {
	//         curr = append(curr, index)
	//         dfs(index+1)
	//     }
	//     left += 1
	//     curr = curr[:pivot]
	//     dfs(index+1)
	// }

	// dfs(0)
	// return ans
}
