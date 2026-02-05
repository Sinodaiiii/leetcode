package easy

import "sort"

func inventoryManagement260205(stock []int, cnt int) []int {
	sort.Ints(stock)
	return stock[:cnt]
}

func inventoryManagement2602052(stock []int, cnt int) []int {
	if cnt == 0 {
		return []int{}
	}
	var findNum func(l, r int)
	findNum = func(l, r int) {
		// fmt.Println(l, r, stock)
		pivot := stock[l]
		next, curr := l+1, l+1
		for curr <= r {
			if stock[curr] <= pivot {
				stock[curr], stock[next] = stock[next], stock[curr]
				next += 1
			}
			curr += 1
		}
		stock[next-1], stock[l] = stock[l], stock[next-1]
		// fmt.Println(next)
		if next == cnt {
			return
		} else if next > cnt {
			findNum(l, next-2)
		} else {
			findNum(next, r)
		}
	}

	findNum(0, len(stock)-1)
	return stock[:cnt]
}
