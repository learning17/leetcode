package main

import (
	"fmt"
)

func main() {
	w := WriteBufferSize(12)
	s := &serverOptions{}
	w.apply(s)
	fmt.Println(s)
}

type serverOptions struct {
	writeBufferSize int
	readBufferSize  int
}

type ServerOption interface {
	apply(*serverOptions)
}

type funcServerOption struct {
	f func(*serverOptions)
}

func (fdo *funcServerOption) apply(do *serverOptions) {
	fdo.f(do)
}

func newFuncServerOption(f func(*serverOptions)) *funcServerOption {
	return &funcServerOption{
		f: f,
	}
}

func WriteBufferSize(s int) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.writeBufferSize = s
	})
}
