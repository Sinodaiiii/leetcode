package medium

func sumFourDivisors260104(nums []int) int {
	ans := 0
	maxNum := 0
	dict := make([]int, 100001)
	for _, num := range nums {
		dict[num] += 1
		maxNum = max(maxNum, num)
	}
	n := maxNum/2 + 1
	list := make([]bool, n+1)
	store := make([]int, n+1)
	curr := 0
	for i := 2; i <= n; i++ {
		if list[i] {
			continue
		}
		store[curr] = i
		curr += 1
		for j := 2; i*j <= n; j++ {
			list[i*j] = true
		}
	}
	for k := 0; k < curr; k++ {
		i := store[k]
		for l := 0; l < k; l++ {
			j := store[l]
			if i*j > maxNum {
				break
			}
			if dict[i*j] > 0 {
				ans += (1 + i + j + i*j) * dict[i*j]
			}
		}
	}
	for _, i := range store {
		if i*i*i > maxNum {
			break
		}
		if dict[i*i*i] > 0 {
			ans += (1 + i + i*i + i*i*i) * dict[i*i*i]
		}
	}
	return ans
}
