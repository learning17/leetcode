package main
// https://leetcode-cn.com/problems/maximum-frequency-stack/
import (
	"container/list"
)

type Node struct {
	Value int
	Freq int
}

type FreqStack struct {
	maxFreq int
	valueDict map[int]*list.Element
	freqDict map[int]*list.List 
}

func Constructor() FreqStack {
	return FreqStack {
		valueDict: make(map[int]*list.Element),
		freqDict: make(map[int]*list.List),
	}
}

func (this *FreqStack) Push(val int)  {
	var nodeP *Node
	if e, ok := this.valueDict[val]; ok {
		nodeP = e.Value.(*Node)
		nodeP.Freq++
	} else {
		nodeP = &Node{val, 1}
	}
	if nodeP.Freq > this.maxFreq {
		this.maxFreq = nodeP.Freq 
	}
	if _, ok := this.freqDict[nodeP.Freq]; !ok {
		this.freqDict[nodeP.Freq] = list.New()
	}
	this.valueDict[val] = this.freqDict[nodeP.Freq].PushFront(nodeP)
}

func (this *FreqStack) Pop() int {
	if this.maxFreq == 0 {
		return -1
	}
	e := this.freqDict[this.maxFreq].Front()
	this.freqDict[this.maxFreq].Remove(e)
	if this.freqDict[this.maxFreq].Len() == 0 {
		this.maxFreq--
	}
	node := e.Value.(*Node)
	node.Freq--
	if node.Freq == 0 {
		delete(this.valueDict, node.Value)
	}
	return node.Value
}


