package main

import (
	"fmt"
	"sync"
)

func main() {
	dogCh := make(chan int, 1)
	catCh := make(chan int, 1)
	fishCh := make(chan int, 1)
	wg := &sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(3)
		go dog(dogCh, catCh, wg)
		go cat(catCh, fishCh, wg)
		go fish(fishCh, dogCh, wg)
	}
	dogCh <- 0
	wg.Wait()
}

func dog(out chan int, in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-out
	fmt.Println("dog")
	in <- 0
	fmt.Println("after dog")
}

func cat(out chan int, in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-out
	fmt.Println("cat")
	in <- 0
	fmt.Println("after cat")
}

func fish(out chan int, in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-out
	fmt.Println("fish")
	in <- 0
	fmt.Println("after fish")
}
