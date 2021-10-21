package main
// https://www.nowcoder.com/practice/05fed41805ae4394ab6607d0d745c8e4

func minEditCost( str1 string ,  str2 string ,  ic int ,  dc int ,  rc int ) int {
	size1, size2 := len(str1), len(str2)
	if size1 == 0 {
		return ic*size2
	}
	if size2 == 0 {
		return dc*size1
	}
	dp := make([][]int, size1+1)
	for i := 0; i < size1+1; i++ {
		dp[i] = make([]int, size2+1)
		dp[i][0] = i*dc
	}
	for j := 0; j < size2; j++ {
		dp[0][j] = j*ic
	}
	for i := 1; i < size1+1; i++ {
		for j := 1; j < size2+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+ic, dp[i-1][j]+dc, dp[i-1][j-1]+rc)
			}
		}
	}
	return dp[size1][size2]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

