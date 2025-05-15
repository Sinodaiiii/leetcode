package medium

func smallestSubsequence(s string) string {
	dict := map[uint8]bool{}
	n := len(s)
	for i:=n-1; i>=0; i-- {
		if dict[s[i]] {
			s = append(s[:i], [i+1:]...)
		}
	}
}
