package main
// https://leetcode-cn.com/problems/implement-stack-using-queues/
// 用队列实现栈

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	size := len(this.queue)
	for i := size -1; i > 0; i-- {
		x := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, x)
	}
	x := this.queue[0]
	this.queue = this.queue[1:]
	return x 
}

func (this *MyStack) Top() int {
	x := this.Pop()
	this.queue = append(this.queue, x)
	return x
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

