package main
// https://www.nowcoder.com/practice/1a3de8b83d12437aa05694b90e02f47a
func reverse( x int ) int {
	sign := 1
	if x < 0 {
		x, sign = -x, -1
	}
	ans := 0
	for ;x > 0; x /= 10 {
		ans = 10*ans + x % 10
	}
	ans = sign*ans
	if ans > (1 << 31)-1 || ans < -(1 << 31) {
		return 0
	}
	return ans
}

