package main
// https://www.nowcoder.com/practice/c6c7742f5ba7442aada113136ddea0c3

func Fibonacci( n int ) int {
	if n < 3 {
		return 1
	}
	dp0, dp1 := 1, 1
	for i := 3; i <= n; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}
	return dp1
}

