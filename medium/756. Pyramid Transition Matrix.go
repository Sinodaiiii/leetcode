package medium

func pyramidTransition(bottom string, allowed []string) bool {
	dict := map[string][]string{}
	for _, s := range allowed {
		dict[s[:2]] = append(dict[s[:2]], s[2:])
	}
	storage := map[string]bool{}
	var checkStage func(str string) bool
	checkStage = func(str string) bool {
		if storage[str] {
			return false
		}
		if len(str) == 1 {
			return true
		}

		nextStrs := []string{}
		for i := 1; i < len(str); i++ {
			curr := []string{}
			if v, exist := dict[str[i-1:i+1]]; exist {
				for _, v2 := range v {
					if len(nextStrs) == 0 {
						curr = append(curr, v2)
					} else {
						for _, nextStr := range nextStrs {
							curr = append(curr, nextStr+v2)
						}
					}
				}
			} else {
				return false
			}
			nextStrs = curr
		}
		for _, nextStr := range nextStrs {
			if checkStage(nextStr) {
				return true
			}
			storage[nextStr] = true
		}
		return false
	}

	return checkStage(bottom)
}
