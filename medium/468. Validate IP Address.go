package medium

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	input := strings.Split(queryIP, ".")
	if len(input) == 4 {
		for _, s := range input {
			num, _ := strconv.Atoi(s)
			if num > 255 || strconv.Itoa(num) != s {
				return "Neither"
			}
		}
		return "IPv4"
	}
	input = strings.Split(queryIP, ":")
	// fmt.Println(input)
	if len(input) == 8 {
		for _, s := range input {
			// fmt.Println(s)
			if len(s) > 4 || len(s) == 0 {
				return "Neither"
			}
			for _, c := range s {
				if (c > 'f' && c <= 'z') || (c > 'F' && c <= 'Z') {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}
	return "Neither"
}
