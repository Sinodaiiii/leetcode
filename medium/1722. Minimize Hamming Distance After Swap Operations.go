package medium

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	father := make([]int, n)
	for i := 0; i < n; i++ {
		father[i] = i
	}

	var find func(index int) int
	var union func(x, y int)

	find = func(index int) int {
		if father[index] != index {
			father[index] = find(father[index])
		}
		return father[index]
	}
	union = func(x, y int) {
		fx, fy := find(x), find(y)
		if fx != fy {
			father[fy] = fx
		}
	}

	for _, pair := range allowedSwaps {
		if find(pair[0]) != find(pair[1]) {
			union(pair[0], pair[1])
		}
		// fmt.Println(pair)
		// fmt.Println(father)
	}
	dict := map[int][]int{}
	for i := 0; i < n; i++ {
		f := find(i)
		dict[f] = append(dict[f], i)
	}
	// fmt.Println(dict)
	ans := 0
	for _, v := range dict {
		currNum := map[int]int{}
		for _, i := range v {
			currNum[source[i]] += 1
			currNum[target[i]] -= 1
		}
		// fmt.Println(currNum)
		for _, i := range currNum {
			if i > 0 {
				ans += i
			}
		}
	}
	return ans
}
