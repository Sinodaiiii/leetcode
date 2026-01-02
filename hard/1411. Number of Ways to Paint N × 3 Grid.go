package hard

func numOfWays260103(n int) int {
	type1, type2 := 6, 6
	for i := 2; i <= n; i++ {
		type1, type2 = (type1*3+type2*2)%1000000007, (type1*2+type2*2)%1000000007
	}
	return (type1 + type2) % 1000000007
}
