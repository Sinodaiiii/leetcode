package hard

func minJumps260518(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}
	dict := map[int][]int{}
	for index, num := range arr {
		dict[num] = append(dict[num], index)
	}
	queue := []int{0}
	visited := make([]bool, n)
	visited[0] = true
	step := 0
	for len(queue) != 0 {
		step += 1
		tmp := make([]int, 0)
		for _, i := range queue {
			for _, index := range []int{i - 1, i + 1} {
				if index >= 0 && index < n && !visited[index] {
					if index == n-1 {
						return step
					}
					visited[index] = true
					tmp = append(tmp, index)
				}
			}
			for _, index := range dict[arr[i]] {
				if index >= 0 && index < n && !visited[index] {
					if index == n-1 {
						return step
					}
					visited[index] = true
					tmp = append(tmp, index)
				}
			}
			delete(dict, arr[i])
		}
		queue = tmp
	}
	return step
}
