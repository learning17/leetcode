package main
// https://leetcode-cn.com/problems/implement-trie-prefix-tree/

type Trie struct {
	children [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _,c := range word {
		index := c - 'a'
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		index := c - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		index := c - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}
