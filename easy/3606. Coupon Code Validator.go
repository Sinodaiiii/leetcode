package easy

import "sort"

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	dict := map[string][]int{}
	curr := true
	for i := 0; i < len(code); i++ {
		curr = true
		if !isActive[i] || len(code[i]) == 0 {
			curr = false
		}
		for _, s := range code[i] {
			if !(s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' || s >= '0' && s <= '9' || s == '_') {
				curr = false
				break
			}
		}
		if !curr {
			continue
		}
		dict[businessLine[i]] = append(dict[businessLine[i]], i)
	}
	ans := []string{}
	check := func(str string) {
		if len(dict[str]) > 0 {
			sort.Slice(dict[str], func(i, j int) bool { return code[dict[str][i]] <= code[dict[str][j]] })
			for _, index := range dict[str] {
				ans = append(ans, code[index])
			}
		}
	}
	check("electronics")
	check("grocery")
	check("pharmacy")
	check("restaurant")

	return ans
}
