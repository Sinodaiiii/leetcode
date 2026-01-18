package goroutine

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Main_goroutine() {
	// 1. 初始化一个 1,000,000 长度的数组（示例：填充 0~999,999）
	arr := make([]int, 1_000_000)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	target := 999999 // 要查找的目标值

	// 2. 使用 context 控制 goroutine 的取消
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 确保所有 goroutine 都能被终止

	// 3. 使用 WaitGroup 等待所有 goroutine 结束
	var wg sync.WaitGroup
	resultChan := make(chan int, 1) // 用于接收结果

	segmentSize := len(arr) / 10 // 每段大小

	// 4. 启动 10 个 goroutine 并发搜索
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(segment int) {
			defer wg.Done()

			start := segment * segmentSize
			end := start + segmentSize
			if segment == 9 { // 最后一段可能稍大
				end = len(arr)
			}

			// 5. 在当前段搜索目标值
			for j := start; j < end; j++ {
				select {
				case <-ctx.Done(): // 如果收到取消信号，立即退出
					return
				default:
					if arr[j] == target {
						// 6. 找到目标值，发送结果并取消其他 goroutine
						select {
						case resultChan <- j:
							cancel() // 通知其他 goroutine 停止
						default:
						}
						return
					}
				}
			}
		}(i)
	}

	// 7. 等待第一个结果或所有 goroutine 结束
	go func() {
		wg.Wait()
		close(resultChan) // 所有 goroutine 结束，关闭 channel
	}()

	// 8. 接收结果
	select {
	case index, ok := <-resultChan:
		if ok {
			fmt.Printf("Found target %d at index: %d\n", target, index)
		} else {
			fmt.Printf("Target %d not found\n", target)
		}
	case <-time.After(5 * time.Second): // 超时保护
		fmt.Println("Search timeout")
	}
}
