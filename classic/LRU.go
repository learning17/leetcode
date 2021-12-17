package main
// https://leetcode-cn.com/problems/OrIXps/

// 基于 container/list 实现 LRU
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
	dict map[int]*list.Element
	linkList *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		dict: make(map[int]*list.Element),
		linkList: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.dict[key]
	if !ok {
		return -1
	}
	this.linkList.MoveToFront(e)
	return e.Value.(*Node).Val
}

func (this *LRUCache) Put(key int, value int)  {
	e, ok := this.dict[key]
	if ok {
		this.linkList.MoveToFront(e)
		node := e.Value.(*Node)
		node.Val = value
		return
	}
	if this.capacity == this.size {
		e = this.linkList.Back()
		this.linkList.Remove(e)
		delete(this.dict, e.Value.(*Node).Key)
		this.size--
	}
	this.size++
	node := &Node{key, value}
	this.dict[key] = this.linkList.PushFront(node)
}

/*
type LinkNode struct {
	Key int
	Value int
	Pre *LinkNode
	Next *LinkNode
} 

type LinkList struct {
	head *LinkNode
	tail *LinkNode
}

func (this *LinkList) MoveToFront(node *LinkNode) {
	if node.Pre == nil {
		return
	}
	if node.Next == nil {
		this.tail = node.Pre
		node.Pre.Next, node.Pre = nil, nil 
	} else {
		node.Pre.Next, node.Next.Pre = node.Next, node.Pre
		node.Pre, node.Next = nil, nil
	}
	node.Next, this.head.Pre = this.head, node
	this.head = node
}

func (this *LinkList) Remove(node *LinkNode) {
	if node.Pre != nil {
		node.Pre.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Pre = node.Pre
	}
	if node.Pre == nil {
		this.head = node.Next
	}
	if node.Next == nil {
		this.tail = node.Pre
	}
	node.Pre, node.Next = nil, nil
}

func (this *LinkList) PushFront(node *LinkNode) {
	if this.head == nil {
		this.head, this.tail = node, node
		return
	}
	node.Next, this.head.Pre = this.head, node
	this.head = node
}

type LRUCache struct {
	capacity int
	size int
	dict map[int]*LinkNode
	linkList *LinkList
}

func Constructor(capacity int) LRUCache {
	return LRUCache {
		capacity: capacity,
		dict: make(map[int]*LinkNode),
		linkList: new(LinkList),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.dict[key]
	if !ok {
		return -1
	}
	this.linkList.MoveToFront(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.dict[key]
	if ok {
		node.Value = value
		this.linkList.MoveToFront(node)
		return
	}
	if this.capacity == this.size {
		node := this.linkList.tail
		delete(this.dict, node.Key)
		this.linkList.Remove(node)
		this.size--
	}
	this.size++
	node = &LinkNode{key, value, nil, nil}
	this.linkList.PushFront(node)
	this.dict[key] = node
}*/
