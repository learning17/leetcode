package main
// https://leetcode-cn.com/problems/lfu-cache/

import (
	"container/list"
)

type Node struct {
	Key int
	Value int
	Freq int
}

type LFUCache struct {
	capacity int
	size int
	minFre int
	keyDict map[int]*list.Element
	freqDict map[int]*list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache {
		capacity: capacity,
		keyDict: make(map[int]*list.Element),
		freqDict: make(map[int]*list.List),
	}
}

func (this *LFUCache) Get(key int) int {
	e, ok := this.keyDict[key]
	if !ok {
		return -1
	}
	this.moveToHighLevel(e)
	return e.Value.(*Node).Value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	e, ok := this.keyDict[key]
	if ok {
		p,_ := e.Value.(*Node)
		p.Value = value
		this.moveToHighLevel(e)
		return
	}
	if this.size == this.capacity {
		linkList, _ := this.freqDict[this.minFre]
		e := linkList.Back()
		delete(this.keyDict, e.Value.(*Node).Key)
		linkList.Remove(e)
		this.size--
	}
	this.size++
	this.minFre = 1
	node := &Node {
		Key: key,
		Value: value,
		Freq: 1,
	}
	if _,ok := this.freqDict[node.Freq]; !ok {
		this.freqDict[node.Freq] = list.New()
	}
	this.keyDict[key] = this.freqDict[node.Freq].PushFront(node)
}

func (this *LFUCache) moveToHighLevel(e *list.Element) {
	node := e.Value.(*Node)
	oldList,_ := this.freqDict[node.Freq]
	oldList.Remove(e)
	if oldList.Len() == 0 && this.minFre == node.Freq {
		this.minFre++
	}
	node.Freq++
	if _, ok := this.freqDict[node.Freq]; !ok {
		this.freqDict[node.Freq] = list.New()
	}
	this.keyDict[node.Key] = this.freqDict[node.Freq].PushFront(node)
}
