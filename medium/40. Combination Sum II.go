package medium

import "sort"

func csum(target int, dict map[int]int, keys []int, index int) [][]int {
	ret := make([][]int, 0)
	for i := index + 1; i < len(keys); i++ {
		for j := 1; j <= dict[keys[i]] && target-j*keys[i] >= 0; j++ {
			if target-j*keys[i] == 0 {
				nums := []int{}
				for k := 0; k < j; k++ {
					nums = append(nums, keys[i])
				}
				ret = append(ret, nums)
			} else {
				nums := csum(target-j*keys[i], dict, keys, i)
				if len(nums) > 0 {
					for k := 0; k < len(nums); k++ {
						num := []int{}
						for l := 0; l < j; l++ {
							num = append(num, keys[i])
						}
						ret = append(ret, append(num, nums[k]...))
					}
				}
			}
		}
	}
	return ret
}

func combinationSum2(candidates []int, target int) [][]int {
	dict := map[int]int{}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	for i := 0; i < len(candidates); i++ {
		dict[candidates[i]] += 1
	}
	keys := []int{}
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	ret := make([][]int, 0)
	for i := 0; i < len(keys); i++ {
		for j := 1; j <= dict[keys[i]] && target-j*keys[i] >= 0; j++ {
			if target-j*keys[i] == 0 {
				nums := []int{}
				for k := 0; k < j; k++ {
					nums = append(nums, keys[i])
				}
				ret = append(ret, nums)
			} else {
				nums := csum(target-j*keys[i], dict, keys, i)
				if len(nums) > 0 {
					for k := 0; k < len(nums); k++ {
						num := []int{}
						for l := 0; l < j; l++ {
							num = append(num, keys[i])
						}
						ret = append(ret, append(num, nums[k]...))
					}
				}
			}
		}
	}
	return ret
}
