package main

import (
	"context"
	"fmt"
)

func main() {
	var h = func(ctx context.Context, req interface{}) (interface{}, error) {
		fmt.Println("do something")
		return "", nil
	}
	var inter1 = func(ctx context.Context, req interface{}, handler UnaryHandler) (interface{}, error) {
		fmt.Println("inter1")
		return handler(ctx, req)
	}
	var inter2 = func(ctx context.Context, req interface{}, handler UnaryHandler) (interface{}, error) {
		fmt.Println("inter2")
		return handler(ctx, req)
	}
	interceptors := []UnaryServerInterceptor{inter1, inter2}
	initInter := chainUnaryInterceptors(interceptors)
	var ctx context.Context
	initInter(ctx, "req", h)
}

type UnaryHandler func(ctx context.Context, req interface{}) (interface{}, error)
type UnaryServerInterceptor func(ctx context.Context, req interface{}, handler UnaryHandler) (interface{}, error)

func chainUnaryInterceptors(interceptors []UnaryServerInterceptor) UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, handler UnaryHandler) (interface{}, error) {
		var state struct {
			i    int
			next UnaryHandler
		}
		state.next = func(ctx context.Context, req interface{}) (interface{}, error) {
			if state.i == len(interceptors)-1 {
				return interceptors[state.i](ctx, req, handler)
			}
			state.i++
			return interceptors[state.i-1](ctx, req, state.next)
		}
		return state.next(ctx, req)
	}
}
