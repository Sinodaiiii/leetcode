package medium

import "math"

func minimumCost260129(source string, target string, original []byte, changed []byte, cost []int) int64 {
	matrix := make([][]int, 26)
	for i := 0; i < 26; i++ {
		matrix[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			matrix[i][j] = math.MaxInt32
		}
	}
	for i := range original {
		matrix[original[i]-'a'][changed[i]-'a'] = min(matrix[original[i]-'a'][changed[i]-'a'], cost[i])
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 26; k++ {
				if matrix[j][i] < math.MaxInt32 && matrix[i][k] < math.MaxInt32 {
					matrix[j][k] = min(matrix[j][k], matrix[j][i]+matrix[i][k])
				}
			}
		}
	}
	// fmt.Println(matrix)
	ans := 0
	for i, c := range source {
		if byte(c) == target[i] {
			continue
		}
		if matrix[c-'a'][target[i]-'a'] == math.MaxInt32 {
			return int64(-1)
		}
		ans += matrix[c-'a'][target[i]-'a']
	}
	return int64(ans)
}
