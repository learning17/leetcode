package main
// https://www.nowcoder.com/practice/4c776177d2c04c2494f2555c9fcc1e49
var stack []int
var minStack []int
func Push(node int) {
	stack = append(stack, node)
	if len(minStack) == 0 || minStack[len(minStack)-1] >= node {
		minStack = append(minStack, node)
	} else {
		minStack = append(minStack, minStack[len(minStack)-1])
	}
}

func Pop() {
	stack = stack[:len(stack)-1]
	minStack = minStack[:len(minStack)-1]
}

func Top() int {
	return stack[len(stack)-1]
}

func Min() int {
	return minStack[len(minStack)-1]
}

