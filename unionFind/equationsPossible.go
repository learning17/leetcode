package main
// https://leetcode-cn.com/problems/satisfiability-of-equality-equations/
func equationsPossible(equations []string) bool {
	if len(equations) == 0 {
		return true
	}
	u := Constructor(26)
	for i := 0; i < len(equations); i++ {
		if equations[i][1] == '!' {
			continue
		}
		u.union(int(equations[i][0]-'a'), int(equations[i][3]-'a'))
	}
	for i := 0; i < len(equations); i++ {
		if equations[i][1] != '!' {
			continue
		}
		if equations[i][0] == equations[i][3] {
			return false
		}
		if u.isConnected(int(equations[i][0]-'a'), int(equations[i][3]-'a')) {
			return false
		}
	}
	return true
}


type UF struct {
	count int
	parent []int
	size []int
}

func Constructor(count int) *UF {
	u := &UF {
		count: count,
		parent: make([]int, count),
		size: make([]int, count),
	}
	for i := 0; i < count; i++ {
		u.parent[i] = i
		u.size[i] = i
	}
	return u
}

func (u *UF) find(x int) int {
	for x != u.parent[x] {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *UF) union(p, q int) {
	pRoot, qRoot := u.find(p), u.find(q)
	if pRoot == qRoot {
		return
	}
	if u.size[pRoot] <= u.size[qRoot] {
		u.parent[pRoot] = qRoot
		u.size[qRoot] += u.size[pRoot]
	} else {
		u.parent[qRoot] = pRoot
		u.size[pRoot] += u.size[qRoot]
	}
}

func (u *UF) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}
