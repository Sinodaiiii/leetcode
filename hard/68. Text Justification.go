package hard

func fullJustify(words []string, maxWidth int) []string {
	n := len(words)
	ans := []string{}
	left := 0
	currLen := len(words[0])
	for i := 1; i < n; i++ {
		if currLen+len(words[i])+1 > maxWidth {
			if i-left == 1 {
				currS := words[left]
				for j := left + 1; j < i; j++ {
					currS += " " + words[j]
				}
				for j := len(currS); j < maxWidth; j++ {
					currS += " "
				}
				ans = append(ans, currS)
				currLen = len(words[i])
				left = i
				continue
			}
			wordLen := 0
			for j := left; j < i; j++ {
				wordLen += len(words[j])
			}
			spaceLen := maxWidth - wordLen
			spaceNum := i - left - 1
			spaceBetween := spaceLen / spaceNum
			spaceAdd := spaceLen % spaceNum
			currS := words[left]
			for j := left + 1; j < i; j++ {
				for k := 0; k < spaceBetween; k++ {
					currS += " "
				}
				if spaceAdd > 0 {
					currS += " "
					spaceAdd -= 1
				}
				currS += words[j]
			}
			ans = append(ans, currS)
			currLen = len(words[i])
			left = i
		} else {
			currLen += len(words[i]) + 1
		}
	}
	currS := words[left]
	for j := left + 1; j < n; j++ {
		currS += " " + words[j]
	}
	for j := len(currS); j < maxWidth; j++ {
		currS += " "
	}
	ans = append(ans, currS)
	return ans
}
