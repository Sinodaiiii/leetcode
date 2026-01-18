package acm_IO

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var lines [10]string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("请输入10行内容：")
	for i := 0; i < 10; i++ {
		if scanner.Scan() {
			lines[i] = scanner.Text()
		} else {
			fmt.Printf("读取第 %d 行时出错\n", i+1)
			return
		}
	}

	fmt.Println("\n输入的内容如下：")
	for i, line := range lines {
		fmt.Printf("第 %d 行：%s\n", i+1, line)
	}
}
