package medium

func longestBeautifulSubstring(word string) int {
	n := len(word)
	ans := 0
	curr := 0
	legal := false
	var preChar uint8
	for i := 0; i < n; i++ {
		switch word[i] {
		case 'a':
			legal = true
			if preChar != 'a' {
				curr = 1
			} else {
				curr++
			}
		case 'e':
			if (preChar == 'a' || preChar == 'e') && legal {
				curr++
			} else {
				legal = false
				curr = 0
			}
		case 'i':
			if (preChar == 'e' || preChar == 'i') && legal {
				curr++
			} else {
				legal = false
				curr = 0
			}
		case 'o':
			if (preChar == 'i' || preChar == 'o') && legal {
				curr++
			} else {
				legal = false
				curr = 0
			}
		case 'u':
			if (preChar == 'o' || preChar == 'u') && legal {
				curr++
				ans = max(ans, curr)
			} else {
				legal = false
				curr = 0
			}
		}
		preChar = word[i]
	}
	return ans
}
