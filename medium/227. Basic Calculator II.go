package medium

func calculate(s string) int {
	numStack := make([]int, 0)
	symbolStack := []int{1}
	getNum := func(begin int) (int, int) {
		num := 0
		for begin < len(s) {
			if s[begin] >= '0' && s[begin] <= '9' {
				num = num*10 + int(s[begin]-'0')
			} else {
				break
			}
			begin += 1
		}
		return num, begin - 1
	}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else if s[i] >= '0' && s[i] <= '9' {
			num := 0
			num, i = getNum(i)
			numStack = append(numStack, num)
		} else if s[i] == '+' {
			symbolStack = append(symbolStack, 1)
		} else if s[i] == '-' {
			symbolStack = append(symbolStack, -1)
		} else {
			j := i + 1
			for j < len(s) && s[j] == ' ' {
				j += 1
			}
			num, newi := getNum(j)
			if s[i] == '*' {
				numStack[len(numStack)-1] = numStack[len(numStack)-1] * num
			} else {
				numStack[len(numStack)-1] = numStack[len(numStack)-1] / num
			}
			i = newi
		}
	}
	ans := 0
	for i := 0; i < len(numStack); i++ {
		ans += numStack[i] * symbolStack[i]
	}
	return ans
}

func calculate260109(s string) int {
	preResult := 0
	flag := 1
	currResult := 0
	currFlag := 0 // 2:*  3:/
	currNum := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			currNum = currNum*10 + int(c-'0')
		} else if c == ' ' {
			continue
		} else {
			if currFlag == 2 {
				currResult *= currNum
			} else if currFlag == 3 {
				currResult /= currNum
			} else {
				currResult = currNum
			}
			if c == '*' || c == '/' {
				if c == '*' {
					currFlag = 2
				} else if c == '/' {
					currFlag = 3
				}
			} else if c == '+' || c == '-' {
				preResult += flag * currResult
				if c == '+' {
					flag = 1
				} else if c == '-' {
					flag = -1
				}
				currFlag = 0
			}
			currNum = 0
		}
	}
	if currFlag == 2 {
		currResult *= currNum
	} else if currFlag == 3 {
		currResult /= currNum
	} else {
		currResult = currNum
	}
	return preResult + flag*currResult
}
