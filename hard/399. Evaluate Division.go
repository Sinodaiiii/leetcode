package hard

import "math"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	dict := map[string]int{}
	i := 0
	for _, value := range equations {
		if _, exist := dict[value[0]]; exist == false {
			dict[value[0]] = i
			i++
		}
		if _, exist := dict[value[1]]; exist == false {
			dict[value[1]] = i
			i++
		}
	}
	matrix := make([][]float64, i)
	for j := range matrix {
		matrix[j] = make([]float64, i)
		for k := 0; k < i; k++ {
			matrix[j][k] = math.MaxFloat64
		}
	}
	for j, v := range equations {
		matrix[dict[v[0]]][dict[v[1]]] = values[j]
		matrix[dict[v[1]]][dict[v[0]]] = 1 / values[j]
	}
	for j := 0; j < i; j++ {
		for k := 0; k < i; k++ {
			if matrix[k][j] != math.MaxFloat64 {
				for l := 0; l < i; l++ {
					if matrix[k][l] == math.MaxFloat64 && matrix[j][l] != math.MaxFloat64 {
						matrix[k][l] = matrix[k][j] * matrix[j][l]
						matrix[l][k] = 1 / matrix[k][l]
					}
				}
			}
		}
	}
	ans := make([]float64, len(queries))
	for j := 0; j < len(queries); j++ {
		_, exist1 := dict[queries[j][0]]
		_, exist2 := dict[queries[j][1]]
		if !exist1 || !exist2 || matrix[dict[queries[j][0]]][dict[queries[j][1]]] == math.MaxFloat64 {
			ans[j] = -1.0
		} else {
			ans[j] = matrix[dict[queries[j][0]]][dict[queries[j][1]]]
		}
	}
	return ans
}
