package goroutine

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const target = 100

func Main_goroutine_printnum_inturn() {
	oddC, evenC := make(chan int, 1), make(chan int, 1)
	fmt.Println("start main proc")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := sync.WaitGroup{}
	print_thread := func(input string) {
		defer wg.Done()
		fmt.Println("start thread " + input)
		inC, outC := oddC, evenC
		if input == "even" {
			inC, outC = evenC, oddC
		}
		for i := 0; i < 1000; i++ {
			select {
			case <-ctx.Done():
				return
			case j := <-inC:
				fmt.Printf("%d ", j)
				if j%10 == 0 {
					fmt.Print("\n")
				}
				if j == target {
					cancel()
					return
				}
				j += 1
				time.Sleep(100 * time.Millisecond)
				outC <- j
			}
		}
	}
	wg.Add(2)
	go print_thread("odd")
	go print_thread("even")
	evenC <- 1
	wg.Wait()
}

func Print3Thread() {
	ch := make([]chan int, 3)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	t := func(index int) {
		defer wg.Done()
		for i := 0; i < 32; i++ {
			select {
			case <-ctx.Done():
				return
			case <-ch[index]:
				fmt.Printf("thread: %d, %d\n", index, i*3+index)
				if index == 0 && i == 20 {
					cancel()
				}
				ch[(index+1)%3] <- 1
			}
		}
	}

	for i := 0; i < 3; i++ {
		ch[i] = make(chan int, 1)
		wg.Add(1)
		go t(i)
	}
	ch[0] <- 1
	wg.Wait()
}
