package medium

func checkInclusion(s1 string, s2 string) bool {
	l, n := len(s1), len(s2)
	if l > n {
		return false
	}
	dict := map[byte]bool{}
	list := make([]int, 26)
	for i := 0; i < l; i++ {
		dict[s1[i]] = true
		list[s1[i]-'a'] += 1
	}
	num := len(dict)
	for i := 0; i < l; i++ {
		if dict[s2[i]] {
			index := int(s2[i] - 'a')
			list[index] -= 1
			if list[index] == 0 {
				num -= 1
			} else if list[index] == -1 {
				num += 1
			}
		}
	}
	for i := l; i < n && num != 0; i++ {
		if dict[s2[i-l]] {
			il := int(s2[i-l] - 'a')
			list[il] += 1
			if list[il] == 0 {
				num -= 1
			} else if list[il] == 1 {
				num += 1
			}
		}
		if dict[s2[i]] {
			ir := int(s2[i] - 'a')
			list[ir] -= 1
			if list[ir] == 0 {
				num -= 1
			} else if list[ir] == -1 {
				num += 1
			}
		}
	}
	if num == 0 {
		return true
	}
	return false
}
