package medium

func addSpaces(s string, spaces []int) string {
	ans := []byte{}
	i := 0
	if spaces[0] == 0 {
		ans = append(ans, ' ')
		i++
	}
	for j := 0; j < len(s); j++ {
		ans = append(ans, s[j])
		if i < len(spaces) && j == spaces[i]-1 {
			i++
			ans = append(ans, ' ')
		}
	}
	return string(ans)
}
