package easy

import "strconv"

func addStrings(num1 string, num2 string) string {
	ans := ""
	flag := 0
	for n1, n2 := len(num1)-1, len(num2)-1; n1 >= 0 || n2 >= 0; {
		a, b := 0, 0
		if n1 >= 0 {
			a = int(num1[n1] - '0')
		}
		if n2 >= 0 {
			b = int(num2[n2] - '0')
		}
		c := a + b + flag
		if c >= 10 {
			flag = 1
		} else {
			flag = 0
		}
		ans = strconv.Itoa(c%10) + ans
		n1--
		n2--
	}
	if flag == 1 {
		ans = "1" + ans
	}
	return ans
}
