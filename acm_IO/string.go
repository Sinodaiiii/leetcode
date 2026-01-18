package acm_IO

import "fmt"

func MainAcmIO() {
	n := 0
	//t := [2]int{}
	//t := [2]int{1, 2}
	input := make([]string, 0)
	//cur := ""
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(n)
	for i := 0; i < n; i++ {
		i1, i2, i3 := "", "", ""
		fmt.Scanln(&i1, &i2, &i3)
		//fmt.Scan(&i3)
		fmt.Println(i1, i2, i3)
	}
	fmt.Print(input)
}
