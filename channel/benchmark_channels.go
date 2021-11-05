package main
import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println(testing.Benchmark(BenchmarkChannelBuffer).String())
}

func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}


func BenchmarkChannelBuffer(b *testing.B) {
	ch := make(chan int,100)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}
