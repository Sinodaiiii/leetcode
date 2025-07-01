package hot100

// len(numstack) <= len(stringStack)
// num'[' == num[numbers]
// add stringStack length when meet '['
// each newString add to the top of stringStack string

func decodeString(s string) string {
	numStack := make([]int, 0)
	curr := 0
	stringStack := []string{}
	for curr < len(s) {
		if s[curr] == '[' {
			stringStack = append(stringStack, "")
		}
		if s[curr] == ']' {
			newString := ""
			for i := 0; i < numStack[len(numStack)-1]; i++ {
				newString += stringStack[len(stringStack)-1]
			}
			numStack = numStack[:len(numStack)-1]
			stringStack = stringStack[:len(stringStack)-1]
			if len(stringStack) > 0 {
				stringStack[len(stringStack)-1] += newString
			} else {
				stringStack = append(stringStack, newString)
			}
		}
		if s[curr] >= '0' && s[curr] <= '9' {
			l := 0
			for i := curr; i < len(s); i++ {
				if s[i] >= '0' && s[i] <= '9' {
					l = l*10 + int(s[i]-'0')
					curr = i
				} else {
					break
				}
			}
			numStack = append(numStack, l)
		}
		if curr < len(s) && s[curr] >= 'a' && s[curr] <= 'z' {
			newString := ""
			for curr < len(s) && s[curr] >= 'a' && s[curr] <= 'z' {
				newString += string(s[curr])
				curr += 1
			}
			if len(stringStack) == 0 {
				stringStack = append(stringStack, newString)
			} else {
				stringStack[len(stringStack)-1] += newString
			}
			continue
		}
		curr += 1
	}
	return stringStack[0]
}
