package main
// https://leetcode-cn.com/problems/sqrtx/
/*
算术平方根
*/
func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right - left)/2
		if mid * mid == x {
			return mid
		} else if mid * mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left-1
}
