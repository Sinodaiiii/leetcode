package easy

func convertToTitle(columnNumber int) string {
	ret := ""
	for columnNumber != 0 {
		ret = string('A'+(columnNumber-1)%26) + ret
		columnNumber = (columnNumber - 1) / 26
	}
	return ret
}
