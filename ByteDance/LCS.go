package main
// https://www.nowcoder.com/practice/6d29638c85bb4ffd80c020fe244baf11

import "fmt"

func main() {
	s1, s2 := "1A2C3D4B56", "B1D23A456A"
	res := LCS(s1, s2)
	fmt.Println(res)
}
func LCS( s1 string ,  s2 string ) string {
	size1, size2 := len(s1), len(s2)
	if size1 == 0 || size2 == 0 {
		return "-1"
	}
	dp := make([][]int, size1+1)
	for i := 0; i < size1+1; i++ {
		dp[i] = make([]int, size2+1)
	}
	for i := 1; i < size1+1; i++ {
		for j := 1; j < size2+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] +1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	var help func(int, int) string
	help = func(i, j int) string {
		if i == 0 || j == 0 {
			return ""
		}
		if dp[i][j] == dp[i-1][j-1] + 1 && s1[i-1] == s2[j-1]{
			return help(i-1, j-1) + string(s1[i-1]) 
		} else if dp[i][j] == dp[i-1][j] {
			return help(i-1, j)
		} else {
			return help(i, j-1)
		}
	}
	ans := help(size1, size2)
	if len(ans) == 0 {
		return "-1"
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
