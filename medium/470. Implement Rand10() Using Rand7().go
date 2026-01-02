package medium

func rand7() int {
	return 0
}

func rand10260102() int {
	for {
		rand1 := (rand7()-1)*7 + rand7() - 1
		if rand1 < 40 {
			return rand1/4 + 1
		} else {
			rand1 = rand1 - 40
		}
		rand2 := (rand7()-1)*7 + rand7() - 1
		if rand2 < 40 {
			return rand2/4 + 1
		} else {
			rand2 = rand2 - 40
		}
		rand3 := rand1*9 + rand2
		if rand3 != 80 {
			return rand3/9 + 1
		}
	}
	return 0
}
