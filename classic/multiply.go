package main
// https://leetcode-cn.com/problems/recursive-mulitply-lcci/

func multiply(A int, B int) int {
	ans := 0
	for i := 0; B != 0; B, i = B >> 1, add(i,1) {
		if B & 1 == 0 {
			continue
		}
		ans = add(ans, A << i)
	}
	return ans
}

func add(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b) << 1
	}
	return a
}
