package main
// https://leetcode-cn.com/problems/lru-cache/
import (
	"container/list"
)


type Node struct {
	Key int
	Val int
}

type LRUCache struct {
	capacity int
	size int
	list *list.List
	dict map[int]*list.Element
}


func Constructor(capacity int) LRUCache {
	return LRUCache {
		capacity: capacity,
		list: list.New(),
		dict: make(map[int]*list.Element),
	}
}


func (this *LRUCache) Get(key int) int {
	e, ok := this.dict[key]
	if !ok {
		return -1
	}
	this.list.MoveToFront(e)
	node, _ := e.Value.(*Node)
	return node.Val 
}


func (this *LRUCache) Put(key int, value int)  {
	e, ok := this.dict[key]
	if ok {
		node, _ := e.Value.(*Node)
		node.Val = value
		this.list.MoveToFront(e)
		return
	}
	if this.capacity == this.size {
		e = this.list.Back()
		node, _ := e.Value.(*Node)
		delete(this.dict, node.Key)
		this.list.Remove(e)
		this.size--
	}
	e = this.list.PushFront(&Node{key, value})
	this.dict[key] = e
	this.size++
}
