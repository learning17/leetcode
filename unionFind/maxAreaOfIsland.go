package main
// https://leetcode-cn.com/problems/ZL6zAn/
// 岛屿的最大面积

func maxAreaOfIsland(grid [][]int) int {
	row_size := len(grid)
	if row_size == 0 {
		return 0
	}
	col_size := len(grid[0])
	u := Constructor(row_size * col_size)
	for r := 1; r < row_size; r++ {
		if grid[r][0] == 1 && grid[r-1][0] == 1 {
			u.union((r-1)*col_size, r*col_size)
		}
		if grid[r][col_size-1] == 1 && grid[r-1][col_size-1] == 1 {
			u.union((r-1)*col_size + col_size-1, r*col_size + col_size-1)
		}
	}
	for c := 1; c < col_size; c++ {
		if grid[0][c] == 1 && grid[0][c-1] == 1 {
			u.union(c-1, c)
		}
		if grid[row_size-1][c] == 1 && grid[row_size-1][c-1] == 1 {
			u.union((row_size-1)*col_size + c-1, (row_size-1)*col_size+ c)
		}
	}
	xDirs := []int{0, -1, 0, 1}
	yDirs := []int{-1, 0, 1, 0}
	for r := 1; r < row_size; r++ {
		for c := 1; c < col_size; c++ {
			if grid[r][c] == 0 {
				continue
			}
			for i := 0; i < 2; i++ {
				rn, cn := r + xDirs[i], c + yDirs[i]
				if grid[rn][cn] == 0 {
					continue
				}
				u.union(r*col_size+c, rn*col_size+cn)
			}
		}
	}

	for r := row_size - 2; r >= 0; r-- {
		for c := col_size - 2; c >= 0; c-- {
			if grid[r][c] == 0 {
				continue
			}
			for i := 2; i < 4; i++ {
				rn, cn := r+xDirs[i], c+yDirs[i]
				if grid[rn][cn] == 0 {
					continue
				}
				u.union(r*col_size+c, rn*col_size+cn)
			}
		}
	}
	dict := make(map[int]int)
	var maxSize int
	for r := 0; r < row_size; r++ {
		for c := 0; c < col_size; c++ {
			if grid[r][c] == 0 {
				continue
			}
			parent := u.find(r*col_size+c)
			dict[parent]++
			if dict[parent] > maxSize {
				maxSize = dict[parent]
			}
		}
	}
	return maxSize
}

type UF struct {
	parent []int
	size []int
}

func Constructor(n int) UF {
	u := UF{
		parent: make([]int, n),
		size: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.size[i] = 1
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
	if u.size[pRoot] >  u.size[qRoot] {
		u.parent[qRoot] = pRoot
		u.size[pRoot] += u.size[qRoot]
	} else {
		u.parent[pRoot] = qRoot
		u.size[qRoot] += u.size[pRoot]
	}
}

func (u *UF) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}
