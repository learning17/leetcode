package main

import (
	"context"
	"fmt"
	"time"
	"sync"
)

func main() {
	Value()
	TimeOut()
	Deadline()
	Cancel()
}
// WithValue
func Value() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Key", "Value")
	GetValue(ctx)
}
func GetValue(ctx context.Context) {
	value := ctx.Value("Key")
	fmt.Println(value)
}
// WithTimeout
func TimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Watch(ctx, wg)
	wg.Wait()
}
func Watch(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("watch timeout", ctx.Err())
			return
		}
	}
}
// WithDeadline
func Deadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Watch(ctx, wg)
	wg.Wait()
} 
//WithCancel 防止协程泄漏
func Cancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range Gen(ctx) {
		fmt.Println(n)
		if n == 1 {
			cancel()
			break
		}
	}
	time.Sleep(10*time.Second)
}
func Gen(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <- ctx.Done():
				fmt.Println("return Gen")
				return
			case ch <- n:
				n++
				time.Sleep(time.Second)
			}
		}
	}()
	return ch
}
