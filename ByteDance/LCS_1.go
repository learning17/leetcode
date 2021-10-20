package main
// https://www.nowcoder.com/practice/f33f5adc55f444baa0e0ca87ad8a6aac

import "fmt"

func main() {
	s1, s2 := "1A2C3D4B56", "B1D23A456A"
	res := LCS(s1, s2)
	fmt.Println(res)
}
func LCS( s1 string ,  s2 string ) string {
	size1, size2 := len(s1), len(s2)
	if size1 == 0 || size2 == 0 {
		return ""
	}
	dp := make([][]int, size1+1)
	for i := 0; i < size1+1; i++ {
		dp[i] = make([]int, size2+1)
	}
	maxSize, pos := 0, 0
	for i := 1; i < size1+1; i++ {
		for j := 1; j < size2+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > maxSize {
				maxSize, pos = dp[i][j], i
			}
		}
	}
	return s1[pos-maxSize:pos]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
