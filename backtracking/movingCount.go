package main
// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
// 机器人的运动范围
func movingCount(m int, n int, k int) int {
	var count int
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	rDirs := []int{0, 0, -1, 1}
	cDirs := []int{-1, 1, 0, 0}
	var backtrack func(r, c int)
	backtrack = func(r, c int) {
		var sum int
		for i := r; i > 0; i /= 10 {
			sum += i % 10
		}
		for i := c; i > 0; i /= 10 {
			sum += i % 10
		}
		visited[r][c] = true
		if sum > k {
			return
		}
		count++
		for i := 0; i < 4; i++ {
			rn, cn := r + rDirs[i], c + cDirs[i]
			if rn < 0 || rn >= m || cn < 0 || cn >= n || visited[rn][cn] {
				continue
			}
			backtrack(rn, cn)
		}
	}
	backtrack(0, 0)
	return count
}
