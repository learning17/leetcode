package main
import (
	"fmt"
)

func main() {
	primes := sieve()
	for prime := range primes {
		fmt.Print(prime, " ")
	}
}

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2;;i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <- in; i%prime != 0{
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	primes := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <- ch
			ch = filter(ch, prime)
			primes <- prime
		}
	}()
	return primes
}
