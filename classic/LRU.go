package main
// https://leetcode-cn.com/problems/OrIXps/

// LRU 基于container 包
import (
	"container/list"
)

type Node struct {
	key int
	value int
}

type LRUCache struct {
	capacity int
	size int
	dict map[int]*list.Element
	linkList *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache {
		capacity: capacity,
		dict: make(map[int]*list.Element),
		linkList: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.dict[key]
	if !ok {
		return -1
	}
	this.linkList.MoveToFront(node)
	return node.Value.(*Node).value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.dict[key]
	if ok {
		p, _ := node.Value.(*Node)
		p.value = value
		this.linkList.MoveToFront(node)
		return
	}
	if this.size == this.capacity {
		node = this.linkList.Back()
		delete(this.dict, node.Value.(*Node).key)
		this.linkList.Remove(node)
		this.size--
	}
	this.size++
	this.dict[key] = this.linkList.PushFront(&Node{key, value})
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
	this.head.Pre, node.Next = node, this.head
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

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.dict[key]
	if ok {
		node.Value = value
		this.linkList.MoveToFront(node)
		return
	}
	if this.size == this.capacity {
		delete(this.dict, this.linkList.tail.Key)
		this.linkList.Remove(this.linkList.tail)
		this.size--
	}
	this.size++
	node = &LinkNode{
		Key: key,
		Value: value,
	}
	this.dict[key] = node
	this.linkList.PushFront(node)
}*/
