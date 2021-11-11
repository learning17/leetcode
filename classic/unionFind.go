package main
// https://leetcode-cn.com/tag/union-find/problemset/

type UF struct {
	count int
	parent []int
	size []int
}

func Constructor(count int) *UF {
	u := &UF{
		count: count,
		parent: make([]int, count),
		size: make([]int, count),
	}
	for i := 0; i < count; i++ {
		u.parent[i] = i
		u.size[i] = 1
	}
	return u
}

func (u *UF) union(p, q int) {
	pRoot, qRoot := u.find(p), u.find(q)
	if pRoot == qRoot {
		return 
	}
	if u.size[pRoot] < u.size[qRoot] {
		u.parent[pRoot] = qRoot
		u.size[qRoot] += u.size[pRoot]
	} else {
		u.parent[qRoot] = pRoot
		u.size[pRoot] += u.size[qRoot]
	}
	u.count--
}

func (u *UF) find(x int) int {
	for x != u.parent[x] {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *UF) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}
