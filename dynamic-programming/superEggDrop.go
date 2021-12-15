package main
// https://leetcode-cn.com/problems/super-egg-drop/
/*
鸡蛋掉落
给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。
已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。
请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少
*/

import (
	"math"
	"fmt"
)

func main() {
	res := superEggDrop(2, 6)
	fmt.Println(res)
}

func superEggDrop(k int, n int) int {
	if k == 1 || n == 0 {
		return n
	}
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, k+1)
		dp[i][1] = i
	}
	for i := 1; i < n+1; i++ {
		for j := 2; j < k+1; j++ {
			dp[i][j] = math.MaxInt64
			for t := 1; t <= i; t++ {
				dp[i][j] = min(dp[i][j], 1+max(dp[t-1][j-1], dp[i-t][j]))
			}
		}
	}
	return dp[n][k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
