package main
import (
	"fmt"
	 "net/http"
	_ "net/http/pprof"
)

func main() {
	ch := make(chan int)
	go generate(ch)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
	for {
		prime := <- ch
		fmt.Print(prime,  " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func generate(ch chan int) {
	for i := 2; ;i++ {
		ch <- i
	}
}

func filter(in chan int, out chan int, prime int) {
	for {
		i := <- in
		if i % prime != 0 {
			out <- i
		}
	}
}

