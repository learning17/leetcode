package main
import (
	"fmt"
)

func main() {
	primes := sieve(100)
	for prime := range primes {
		fmt.Print(prime, " ")
	}
}

func generate(size int) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < size; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	ch := make(chan int)
	go func() {
		for i := range in {
			if i % prime != 0 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func sieve(size int) chan int {
	primes := make(chan int)
	go func() {
		ch := generate(size)
		for {
			prime, ok := <- ch
			if !ok {
				break
			}
			ch = filter(ch, prime)
			primes <- prime
		}
		close(primes)
	}()
	return primes
}
