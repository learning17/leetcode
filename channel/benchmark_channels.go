package main
import (
	"fmt"
	"testing"
	"bytes"
	"strings"
)

func main() {
	fmt.Println("Buffer:", testing.Benchmark(benchMarkAddStringWithBuffer).String())
	fmt.Println("Op:", testing.Benchmark(benchMarkAddStringWithOp).String())
	fmt.Println("Sprintf:", testing.Benchmark(benchMarkAddStringWithSprint).String())
	fmt.Println("Join:", testing.Benchmark(benchMarkAddStringWithJoin).String())
}

func benchMarkAddStringWithBuffer(b *testing.B) {
	hello, world := "hello", "world"
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString(world)
		_ = buffer.String()
	}
}

func benchMarkAddStringWithOp(b *testing.B) {
	hello, world := "hello", "world"
	for i := 0; i < b.N; i++ {
		_ = hello + world 
	}
}


func benchMarkAddStringWithSprint(b *testing.B) {
	hello, world := "hello", "world"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s", hello, world)
	}
}


func benchMarkAddStringWithJoin(b *testing.B) {
	hello, world := "hello", "world"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{hello, world},"")
	}
}
