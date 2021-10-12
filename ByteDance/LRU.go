package main

// https://www.nowcoder.com/practice/e3769a5f49894d49b871c09cadd13a61

type LinkNode struct {
	Val int
	Key int
	Pre *LinkNode
	Next *LinkNode
}

type LRUCache struct {
	capacity int
	size int
	head *LinkNode
	tail *LinkNode
	dict map[int]*LinkNode
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		dict: make(map[int]*LinkNode),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.dict[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key, value int) {
	node, ok := this.dict[key]
	if ok {
		node.Val = value
		this.moveToHead(node)
		return
	}
	node = &LinkNode{
		Key: key,
		Val: value,
	}
	this.dict[key] = node
	if this.size == this.capacity {
		delete(this.dict, this.tail.Key)
		if this.tail.Pre != nil {
			this.tail.Pre.Next = nil
		}
		this.tail, this.tail.Pre = this.tail.Pre, nil
		this.size--
	}
	if this.tail == nil {
		this.head, this.tail = node, node
	} else {
		node.Pre, node.Next = nil, this.head
		this.head, this.head.Pre = node, node
	}
	this.size++
}

func (this *LRUCache) moveToHead(node *LinkNode) {
	if node == this.head {
		return
	}
	if node == this.tail {
		node.Pre.Next, this.tail = nil, node.Pre
	} else {
		node.Pre.Next, node.Next.Pre = node.Next, node.Pre
	}
	node.Pre, node.Next = nil, this.head
	this.head, this.head.Pre = node, node
}

func LRU( operators [][]int ,  k int ) []int {
	lru := Constructor(k)
	var ans []int
	for i := 0; i < len(operators); i++ {
		if operators[i][0] == 1 {
			lru.Put(operators[i][1], operators[i][2])
		} else {
			ans = append(ans, lru.Get(operators[i][1]))
		}
	}
	return ans
}
