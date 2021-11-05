package main
import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go produce(ch)
	go consume(ch, done)
	<-done
}

func produce(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch <-chan int, done chan<- bool) {
	for i := range ch {
		fmt.Println(i)
	}
	done <- true
}
