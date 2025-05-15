package medium

import "fmt"

func decodeString(s string) string {
	numStack := make([]int, 0)
	dict := map[string]string{}
	for i, v := range s {
		if v == '[' {
			num := 0
			multiple := 1
			for j := i - 1; j >= 0; j-- {
				if s[j] >= '0' && s[j] <= '9' {
					num += multiple * int(s[j]-'0')
					multiple *= 10
				} else {
					break
				}
			}
			numStack = append(numStack, num, i)
		} else if v == ']' {
			j := numStack[len(numStack)-1]
			dup := numStack[len(numStack)-2]
			str := s[j : i+1]
			dupStr := ""
			for k := 0; k < dup; k++ {
				dupStr += str[1 : len(str)-1]
			}
			dict[fmt.Sprintf("%d", dup)+str] = dupStr
			numStack = numStack[:len(numStack)-2]
		}
	}
	ans := ""
	for i := 0; i < len(s); {
		if s[i] >= '0' && s[i] <= '9' {
			count := 0
			for j := i; j <= len(s); j++ {
				if s[j] >= '0' && s[j] <= '9' {
					continue
				}
				if s[j] == '[' {
					count++
				}
				if s[j] == ']' {
					count--
				}
				if count == 0 {
					str := s[i : j+1]
					ans = ans + dict[str]
					i = j + 1
					break
				}
			}
		} else {
			ans = ans + string(s[i])
			i++
		}
	}
	return ans
}
