package main
// https://www.nowcoder.com/practice/735a34ff4672498b95660f43b7fcd628
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484751&idx=1&sn=a873c1f51d601bac17f5078c408cc3f6&scene=21#wechat_redirect
import (
	"sort"
)

func miniSpanningTree( n int ,  m int ,  cost [][]int ) int {
	sort.Slice(cost, func(i, j int) bool {return cost[i][2] < cost[j][2]})
	uf := Constructor(n)
	minCost := 0
	for i := 0; i < m; i++ {
		if uf.isCompleted() {
			break
		}
		if uf.isConnected(cost[i][0], cost[i][1]) {
			continue
		}
		minCost += cost[i][2]
		uf.union(cost[i][0], cost[i][1])
	}
	return minCost
}

type UF struct {
	parent []int
	size []int
	count int
}

func Constructor(n int) *UF {
	uf := UF {
		parent: make([]int, n),
		size: make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return &uf
}

func (u *UF) find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *UF) union(p, q int) {
	rootP, rootQ := u.find(p), u.find(q)
	if rootP == rootQ {
		return
	}
	if u.size[rootP] > u.size[rootQ] {
		u.parent[rootQ] = rootP
		u.size[rootP] += u.size[rootQ]
	} else {
		u.parent[rootP] = rootQ
		u.size[rootQ] += u.size[rootP]
	}
	u.count--
}

func (u *UF) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UF) isCompleted() bool {
	return u.count == 0
}
