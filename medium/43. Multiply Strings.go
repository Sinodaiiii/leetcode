package medium

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	preResult := "0"
	for i := n1 - 1; i >= 0; i-- {
		pre := 0
		currResult := preResult[len(preResult)-(n1-1-i):]
		// fmt.Println("loop: ", preResult, currResult, len(preResult)-(n1-1-i), len(preResult), n1-1-i)
		curr1 := int(num1[i] - '0')
		for j := n2 - 1; j >= 0; j-- {
			preNum := 0
			if len(preResult)-(n1-i)-(n2-j)+1 >= 0 {
				// fmt.Println("preIndex: ", len(preResult)-(n1-i)-(n2-j)+1, len(preResult), n1-i, n2-j)
				preNum = int(preResult[len(preResult)-(n1-i)-(n2-j)+1] - '0')
			}
			curr2 := int(num2[j]-'0')*curr1 + pre + preNum
			// fmt.Println(curr2, int(num2[j] - '0'), curr1, pre, preNum, currResult)
			currResult = string(byte(curr2%10+'0')) + currResult
			pre = curr2 / 10
			// fmt.Println(currResult)
		}
		for pre != 0 {
			currResult = string(byte(pre%10+'0')) + currResult
			pre = pre / 10
		}
		preResult = currResult
	}
	return preResult
}
