package main
import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch <- 2
	go f1(ch)
}

func f1(in chan int) {
	fmt.Println(<-in)
}
