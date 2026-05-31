package medium

import "sort"

func asteroidsDestroyed260531(mass int, asteroids []int) bool {
	sort.Slice(asteroids, func(i, j int) bool { return asteroids[i] <= asteroids[j] })
	for _, a := range asteroids {
		if mass >= a {
			mass += a
		} else {
			return false
		}
	}
	return true
}
