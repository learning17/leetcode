package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
    go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func getData(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
