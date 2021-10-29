package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string, 1)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(ch, i, wg)
	}
	for i := 0; i < 50; i++ {
		go consume(ch, i)
	}
	wg.Wait()
}


func produce(ch chan string, i int, wg *sync.WaitGroup) {
	for j := 0; j < 100; j++ {
		ch <- fmt.Sprintf("producer:%d, seq:%d", i, j)
	}
	wg.Done()
} 

func consume(ch chan string, i int) {
	for data := range ch {
		fmt.Printf("consumer:%d, %s\n", i, data)
	}
}
