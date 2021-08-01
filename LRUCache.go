package main

// https://leetcode-cn.com/problems/lru-cache/

type LRUCache struct {
	capacity int
	size int
	head *LinkNode
	tail *LinkNode
	dict map[int]*LinkNode
}

type LinkNode struct {
	key int
	value int
	prev *LinkNode
	next *LinkNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		dict: make(map[int]*LinkNode),
	}
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.dict[key]
	if !ok {
		return -1
	}
	this.move(node)
	return node.value
}

func (this *LRUCache) move(node *LinkNode) {
	if node == this.tail {
		return
	}
	if node == this.head {
		this.head = node.next
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev,node.next, this.tail.next = this.tail, nil, node
	this.tail = node
}

func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.dict[key]
	if ok {
		this.move(node)
		node.value = value
		return
	}
	newNode := &LinkNode{key: key, value: value}
	if this.size == 0 {
		this.head, this.tail = newNode, newNode
	} else if this.size < this.capacity {
		newNode.prev, this.tail.next = this.tail, newNode
		this.tail = newNode
		this.size++
	} else {
		head, next := this.head, this.head.next
		delete(this.dict, head.key)
		if next == nil {
			this.head, this.tail = newNode, newNode
		} else {
			head.next, next.prev = nil, nil
			this.head = next
		}
		newNode.prev, this.tail.next = this.tail, newNode
		this.tail = newNode
		this.size++
	}
	this.dict[key] = newNode
}
