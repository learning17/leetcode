package main
import (
	"context"
	"sync"
	"time"
	"fmt"
	"runtime"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var items []int
	for i := 0; i < 1001; i++ {
		items = append(items, i)
	}
	Handle(ctx, items)
}

func Handle(ctx context.Context, items []int) {
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
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
}

func GetBatchItems(ctx context.Context, wg *sync.WaitGroup, items []int) ([]int,error) {
	defer wg.Done()
	done := make(chan error, 1)
	ch := make(chan []int, 1)
	paincChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				paincChan <- p
			}
		}()
		rsp, err := RcpAction(items)
		ch <- rsp
		done <- err
	}()
	select {
	case p :=<- paincChan:
		panic(p)
	case err :=<-done:
		return <-ch, err
	case <-ctx.Done():
		return []int{},ctx.Err()
	}
}

func RcpAction(items []int)([]int, error) {
	fmt.Println("RcpAction")
	time.Sleep(10*time.Second)
	return []int{}, nil
}
