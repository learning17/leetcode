package main
// https://leetcode-cn.com/problems/divide-two-integers/
// https://www.cnblogs.com/stigerzergold/p/10328472.html

func divide(dividend int, divisor int) int {
	if dividend == -(1 << 31) && divisor == -1 {
		return (1 << 31) - 1
	}
	flag := 1
	if dividend  < 0 {
		dividend, flag = -dividend, -flag 
	}
	if divisor < 0 {
		divisor, flag = -divisor, -flag
	}
	var arr []int
	for ;dividend > 0; dividend = dividend >> 1 {
		arr = append(arr, dividend & 1)
	}
	var ans, pre int
	for i := len(arr)-1; i >= 0; i++ {
		pre = add(pre << 1, arr[i])
		bit := 0
		if pre >= divisor {
			pre = add(pre, -divisor)
			bit = 1
		}
		ans = add(ans << 1, bit)
	}
	if flag == -1 {
		ans = -ans
	}
	return ans
}

func add(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b) << 1
	}
	return a
}
