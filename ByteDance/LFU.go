package main
// https://www.nowcoder.com/practice/93aacb4a887b46d897b00823f30bfea1


func LFU( operators [][]int ,  k int ) []int {
	l := Constructor(k)
	var res []int
	for i := 0; i < len(operators); i++ {
		if operators[i][0] == 1 {
			l.Put(operators[i][1], operators[i][2])
		} else {
			res = append(res, l.Get(operators[i][1]))
		}
	}
	return res
}

type LinkNode struct {
	Key int
	Value int
	Fre int
	Pre *LinkNode
	Next *LinkNode
}

type List struct {
	head *LinkNode
	tail *LinkNode
}

type LFUCache struct {
	capacity int
	size int
	minFreq int
	keyDict map[int] *LinkNode
	freqDict map[int] *List
}

func Constructor(capacity int) LFUCache {
	return LFUCache {
		capacity: capacity,
		keyDict : make(map[int]*LinkNode),
		freqDict: make(map[int]*List),
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.keyDict[key]
	if !ok {
		return -1
	}
	this.moveNode(node)
	return node.Value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	node, ok := this.keyDict[key]
	if ok {
		node.Value = value
		this.moveNode(node)
		return
	}
	node = &LinkNode{
		Key: key,
		Value: value,
		Fre: 1,
	}
	if this.size == this.capacity {
		tmp, _ := this.freqDict[this.minFreq]
		delete(this.keyDict, tmp.tail.Key)
		if tmp.tail.Pre == nil {
			delete(this.freqDict, this.minFreq)
		} else {
			tmp.tail, tmp.tail.Pre.Next, tmp.tail.Pre = tmp.tail.Pre, nil, nil 
		}
		this.size--
	}
	this.size++
	this.minFreq = 1
	this.keyDict[key] = node
	tmp, ok := this.freqDict[this.minFreq]
	if ok {
		tmp.head.Pre, node.Next = node, tmp.head
	} else {
		this.freqDict[this.minFreq] = &List{head: node, tail:node}
	}
	this.freqDict[this.minFreq].head = node
}

func (this *LFUCache) moveNode(node *LinkNode) {
	if node.Pre == nil && node.Next == nil {
		delete(this.freqDict, node.Fre)
		if this.minFreq == node.Fre {
			this.minFreq++
		}
	} else if node.Pre == nil {
		this.freqDict[node.Fre].head = node.Next
		node.Next.Pre, node.Next = nil, nil
	} else if node.Next == nil {
		this.freqDict[node.Fre].tail = node.Pre
		node.Pre.Next, node.Pre = nil, nil
	} else {
		node.Pre.Next, node.Next.Pre = node.Next, node.Pre
		node.Next, node.Pre = nil, nil
	}
	node.Fre++
	tmp, ok := this.freqDict[node.Fre] 
	if ok {
		node.Next, tmp.head.Pre = tmp.head, node
	} else {
		this.freqDict[node.Fre] = &List{head: node, tail: node}
	}
	this.freqDict[node.Fre].head = node
}

