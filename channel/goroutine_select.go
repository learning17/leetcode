package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	suck(ch1, ch2)
}

func pump1(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func pump2(ch chan int) {
	for i := 100; i < 200; i++ {
		ch <- i
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select{
		case v := <- ch1:
			fmt.Println("channel1:", v)
		case v := <- ch2:
			fmt.Println("channel2:", v)
		default:
			return
		}
	}
}
