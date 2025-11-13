package main

func maxOperations(s string) int {
	ans := 0
	pre := 0
	for i, c := range s {
		if c == '0' && i >= 1 && s[i-1] == '1' {
			ans += pre
		} else if c == '1' {
			pre += 1
		}
	}
	return ans
}
