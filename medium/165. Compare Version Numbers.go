package medium

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	var s, l []string
	flag := 1
	if len(v1) < len(v2) {
		s, l = v1, v2
	} else {
		s, l = v2, v1
		flag = -1
	}
	for i := 0; i < len(s); i++ {
		a, _ := strconv.Atoi(s[i])
		b, _ := strconv.Atoi(l[i])
		if a < b {
			return -1 * flag
		} else if a > b {
			return flag
		}
	}
	for i := len(s); i < len(l); i++ {
		a, _ := strconv.Atoi(l[i])
		if a != 0 {
			return -1 * flag
		}
	}
	return 0
}
