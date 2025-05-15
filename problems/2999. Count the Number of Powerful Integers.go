package problems

import "math"

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	n := len(s)
	num := 0
	for i := 0; i < n; i++ {
		num = num*10 + int(s[i]-'0')
	}
	if int64(num) >= finish {
		return 0
	}
	n1, n2 := 0, 0
	tmp := start
	for tmp != 0 {
		n1++
		tmp = tmp / 10
	}
	tmp = finish
	for tmp != 0 {
		n2++
		tmp = tmp / 10
	}
	t1, t2 := 0, 0
	if int64(num) <= start {
		a1, a2 := 0, 0
		if n1 != 1 {
			a1 = int(start) / int(math.Pow(10, float64(n1-1)))
			a2 = int(start) % int(math.Pow(10, float64(n1-1)))
			if a2 < num {
				a1--
			}
		} else {
			a1 = int(start)
		}
		t1 = 1 + (n1-n-1)*limit + min(a1, limit)
	}
	a1, a2 := 0, 0
	if n2 != 1 {
		a1 = int(finish) / int(math.Pow(10, float64(n2-1)))
		a2 = int(finish) % int(math.Pow(10, float64(n2-1)))
		if a2 < num {
			a1--
		}
	} else {
		a1 = int(finish)
	}
	t2 = 1 + (n2-n-1)*limit + min(a1, limit)
	return int64(t2 - t1)
}
