package medium

func numberOfSpecialChars260527(word string) int {
	n := len(word)
	dict := make([]int, 26)
	for i := n - 1; i >= 0; i-- {
		index := int(word[i] - 'a')
		if index >= 0 && index <= 25 {
			if dict[index] == 1 || dict[index] == 2 {
				dict[index] = 2
			} else {
				dict[index] = 3
			}
		} else {
			index = int(word[i] - 'A')
			if dict[index] <= 1 {
				dict[index] = 1
			} else {
				dict[index] = 3
			}
		}
	}
	ans := 0
	for _, state := range dict {
		if state == 2 {
			ans += 1
		}
	}
	return ans
}
