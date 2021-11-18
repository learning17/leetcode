package main
// https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
// 包含min函数的栈

type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
	} else {
		if this.minStack[len(this.minStack)-1] > x {
			this.minStack = append(this.minStack, x)
		} else {
			this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
		}
	}
}

func (this *MinStack) Pop()  {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.minStack) == 0 {
		return -1
	}
	return this.minStack[len(this.minStack)-1]
}
