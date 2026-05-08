package medium

func minJumps(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	maxNum := 1
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	spf := make([]int, maxNum+1)
	isPrime := make([]bool, maxNum+1)
	for i := 2; i <= maxNum; i++ {
		spf[i] = i
		isPrime[i] = true
	}
	for i := 2; i*i <= maxNum; i++ {
		if isPrime[i] {
			for j := i * i; j <= maxNum; j += i {
				isPrime[j] = false
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}

	dict := make([][]int, maxNum+1)
	for i, num := range nums {
		temp := num
		for temp > 1 {
			p := spf[temp]
			if len(dict[p]) == 0 || dict[p][len(dict[p])-1] != i {
				dict[p] = append(dict[p], i)
			}
			for temp%p == 0 {
				temp /= p
			}
		}
	}

	queue := []int{0}
	visited := make([]bool, n)
	visited[0] = true
	step := 0

	for len(queue) > 0 {
		nextQueue := []int{}
		for _, index := range queue {
			if index == n-1 {
				return step
			}

			if index-1 >= 0 && !visited[index-1] {
				nextQueue = append(nextQueue, index-1)
				visited[index-1] = true
			}
			if index+1 < n && !visited[index+1] {
				nextQueue = append(nextQueue, index+1)
				visited[index+1] = true
			}
			if isPrime[nums[index]] {
				p := nums[index]
				jump := dict[p]
				for _, jumpIndex := range jump {
					if !visited[jumpIndex] {
						nextQueue = append(nextQueue, jumpIndex)
						visited[jumpIndex] = true
					}
				}
				dict[p] = nil
			}
		}
		queue = nextQueue
		step++
	}
	return -1
}
