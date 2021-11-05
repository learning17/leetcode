package main
import (
	"fmt"
)

func main() {
	adder := startServer(func(a, b int) int {return a+b})
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i+N
		req.replyc = make(chan int)
		adder <- req
	}
	for i := N-1; i >= 0; i-- {
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at:", i)
		} else {
			fmt.Println("Request ", i, "is ok")
		}
	}
}

type Request struct {
	a int
	b int
	replyc chan int
}

type binOP func(a, b int) int

func run(op binOP, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOP, service chan *Request) {
	for {
		req := <- service
		go run(op, req)
	}
}

func startServer(op binOP) chan *Request {
	reqChan := make(chan *Request)
	go server(op, reqChan)
	return reqChan
}
