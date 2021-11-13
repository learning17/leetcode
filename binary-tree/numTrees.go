package main
// https://leetcode-cn.com/problems/unique-binary-search-trees/

func numTrees(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}
	var count func(int, int) int
	count = func (start, end int) int {
		if start >= end {
			return 1
		}
		if dp[start][end] != 0 {
			return dp[start][end]
		}
		sum := 0
		for i := start; i <= end; i++ {
			left := count(start, i-1)
			right := count(i+1, end)
			sum += left*right
		}
		dp[start][end] = sum
		return sum
	}
	return count(1, n)
}

