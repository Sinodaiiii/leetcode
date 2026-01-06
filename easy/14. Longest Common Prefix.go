package easy

import "math"

func longestCommonPrefix260107(strs []string) string {
	minN := math.MaxInt32
	for _, str := range strs {
		minN = min(minN, len(str))
	}
	for i := 0; i < minN; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[j-1][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minN]
}
