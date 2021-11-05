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
	go monitor(ch, wg)

	rg := &sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		rg.Add(1)
		go consume(ch, i, rg)
	}
	rg.Wait()
}

func monitor(ch chan string, wg *sync.WaitGroup) {
	wg.Wait()
	close(ch)
}

func produce(ch chan string, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 100; j++ {
		ch <- fmt.Sprintf("producer:%d, seq:%d", i, j)
	}
} 

func consume(ch chan string, i int, rg *sync.WaitGroup) {
	defer rg.Done()
	for data := range ch {
		fmt.Printf("consumer:%d, %s\n", i, data)
	}
}
