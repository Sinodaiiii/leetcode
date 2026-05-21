package medium

func longestCommonPrefix260521(arr1 []int, arr2 []int) int {
	type prefix struct {
		length int
		next   []*prefix
	}
	dummy := &prefix{0, make([]*prefix, 10)}
	for _, num := range arr1 {
		mask := 1
		for num/mask != 0 {
			mask *= 10
		}
		mask = mask / 10
		curr := dummy
		for mask != 0 {
			n := num / mask
			num = num % mask
			mask = mask / 10
			if curr.next[n] == nil {
				curr.next[n] = &prefix{curr.length + 1, make([]*prefix, 10)}
			}
			curr = curr.next[n]
			// fmt.Println(curr)
		}
	}
	ans := 0
	for _, num := range arr2 {
		mask := 1
		for num/mask != 0 {
			mask *= 10
		}
		mask = mask / 10
		curr := dummy
		currLength := 0
		for mask != 0 {
			n := num / mask
			num = num % mask
			mask = mask / 10
			if curr.next[n] == nil {
				break
			}
			curr = curr.next[n]
			currLength = curr.length
		}
		ans = max(ans, currLength)
	}
	return ans
}
