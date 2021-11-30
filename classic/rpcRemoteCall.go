package main
import (
	"context"
	"sync"
	"time"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var items []int
	for i := 0; i < 1001; i++ {
		items = append(items, i)
	}
	Handle(ctx, items)
}

func Handle(ctx context.Context, items []int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var i, last int
	wg := &sync.WaitGroup{}
	for i = 1; i <= len(items); i++ {
		if i % 50 == 0 {
			wg.Add(1)
			last = i
			go GetBatchItems(ctx, wg, items[i-50:i])
		}
	}
	if last < len(items) {
		wg.Add(1)
		go GetBatchItems(ctx, wg, items[last:])
	}
	wg.Wait()
}
func GetBatchItems(ctx context.Context, wg *sync.WaitGroup, items []int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			_, err := RcpAction(items)
			if err != nil {
				panic(err)
			}
			return
		}
	}
}

func RcpAction(items []int)([]int, error) {
	fmt.Println("RcpAction")
	time.Sleep(time.Second)
	return []int{}, nil
}
