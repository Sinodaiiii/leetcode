package Template

// interval: [l1, r1], [l2, r2]
func IsOverlap(l1, r1, l2, r2 int) bool {
	if min(r1, r2)-max(l1, l2) > 0 {
		return true
	}
	return false
}
