package main
// https://www.nowcoder.com/practice/8c82a5b80378478f9484d87d1c5f12a4

func jumpFloor( number int ) int {
	if number < 3 {
		return number
	}
	dp0, dp1 := 1, 2
	for i := 3; i <= number; i++ {
		dp1, dp0 = dp0+dp1, dp1
	}
	return dp1
}

