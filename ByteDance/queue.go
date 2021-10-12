package main

var stack1 [] int
var stack2 [] int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int{
	size2 := len(stack2)
	if size2 == 0 {
		size1 := len(stack1)
		stack2 = make([]int, size1)
		for i := size1-1; i >= 0; i-- {
			stack2[size1-i-1] = stack1[i]
		}
		size2 = size1
		stack1 = stack1[:0] 
	}
	value := stack2[size2-1]
	stack2 = stack2[:size2-1]
	return value
}
