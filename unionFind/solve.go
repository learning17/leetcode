package main
// https://leetcode-cn.com/problems/surrounded-regions/

func solve(board [][]byte)  {
	row_size := len(board)
	if row_size == 0 {
		return
	}
	col_size := len(board[0])
	rootPos := col_size * row_size
	u := Constructor(rootPos +1)
	for i := 0; i < row_size; i++ {
		if board[i][0] == 'O' {
			u.union(i*col_size, rootPos)
		}
		if board[i][col_size-1] == 'O' {
			u.union((i+1)*col_size-1, rootPos)
		}
	}

	for i := 0; i < col_size; i++ {
		if board[0][i] == 'O' {
			u.union(i, rootPos)
		}
		if board[row_size-1][i] == 'O' {
			u.union((row_size-1)*col_size+i, rootPos)
		}
	}
	r_dirs := []int{-1, 0, 1, 0}
	c_dirs := []int{0, -1, 0, 1}
	for r := 1; r < row_size-1; r++ {
		for c := 1; c < col_size-1; c++ {
			if board[r][c] == 'X' {
				continue
			}
			for i := 0; i < 2; i++ {
				r_n, c_n := r+r_dirs[i], c+c_dirs[i]
				if board[r_n][c_n] == 'O' {
					u.union(r*col_size+c, r_n*col_size+c_n)
				}
			}
		}
	}
	for r := row_size - 2; r >= 1; r-- {
		for c := col_size - 2; c >= 1; c-- {
			if board[r][c] == 'X' {
				continue
			}
			for i := 2; i < 4; i++ {
				r_n, c_n := r+r_dirs[i], c+c_dirs[i]
				if board[r_n][c_n] == 'O' {
					u.union(r*col_size+c, r_n*col_size+c_n)
				}
			}
		}
	}
	for r := 1; r < row_size-1; r++ {
		for c := 1; c < col_size-1; c++ {
			if board[r][c] == 'X' {
				continue
			}
			if u.isConnected(rootPos, r*col_size+c) {
				continue
			}
			board[r][c] = 'X'
		}
	}
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
	if u.size[pRoot] <= u.size[qRoot] {
		u.parent[pRoot] = qRoot
		u.size[qRoot] += u.size[pRoot]
	} else {
		u.parent[qRoot] = pRoot
		u.size[pRoot] += u.size[qRoot]
	}
	u.count--
}

func (u *UF) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}
