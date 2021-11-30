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
	ch := make(chan []int)
	for i = 1; i <= len(items); i++ {
		if i % 50 == 0 {
			wg.Add(1)
			last = i
			go GetBatchItems(ctx, wg, items[i-50:i], ch)
		}
	}
	if last < len(items) {
		wg.Add(1)
		go GetBatchItems(ctx, wg, items[last:], ch)
	}
	go monitor(wg, ch)
	for c := range ch {
		fmt.Println(c)
	}
}
func monitor(wg *sync.WaitGroup, ch chan []int) {
	wg.Wait()
	close(ch)
}
func GetBatchItems(ctx context.Context, wg *sync.WaitGroup, items []int, ch chan []int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			res, err := RcpAction(items)
			if err != nil {
				panic(err)
			}
			ch <- res
			return
		}
	}
}

func RcpAction(items []int)([]int, error) {
	fmt.Println("RcpAction")
	time.Sleep(time.Second)
	return []int{}, nil
}
