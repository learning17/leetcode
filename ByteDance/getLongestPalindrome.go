package main
// https://www.nowcoder.com/practice/b4525d1d84934cf280439aeecc36f4af

func getLongestPalindrome( A string ,  n int ) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	maxSize := 1
	for l := 1; l < n; l++ {
		for i := 0; i < n; i++ {
			j := i + l
			if j >= n {
				break
			}
			if A[i] == A[j] && dp[i+1][j-1] == j-i-1 {
				dp[i][j] = dp[i+1][j-1] + 2
			}
			maxSize = max(maxSize, dp[i][j])
		}
	}
	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
