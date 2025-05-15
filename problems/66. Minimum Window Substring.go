package problems

import "math"

func minWindow(s string, t string) string {
	dict := map[int32]int{}
	for _, value := range t {
		dict[value] += 1
	}
	ans, anslen := 0, math.MaxInt32
	n := len(t)
	l, r := 0, 0
	for ; l < len(s); l++ {
		if _, exist := dict[int32(s[l])]; exist {
			dict[int32(s[l])]--
			n--
			if n == 0 {
				if r-l+1 < anslen {
					anslen = r - l + 1
					ans = l
				}
			}
			r = l + 1
			break
		}
	}
	for l < len(s) && r < len(s) {
		if num, exist := dict[int32(s[r])]; exist {
			dict[int32(s[r])]--
			if num > 0 {
				n--
				for {
					if n == 0 {
						if r-l+1 < anslen {
							anslen = 1
							ans = l
						}
						dict[int32(s[l])]++
						if dict[int32(s[l])] == 1 {
							n++
						}
						l++
						for ; l < len(s); l++ {
							if _, lexist := dict[int32(s[l])]; lexist {
								break
							}
						}
					} else {
						break
					}
				}
			}
		}
		r++
	}
	if anslen == math.MaxInt32 {
		return ""
	}
	return s[ans : ans+anslen]
}
