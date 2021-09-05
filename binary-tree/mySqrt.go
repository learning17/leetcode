package main
// https://leetcode-cn.com/problems/sqrtx/

func mySqrt(x int) int {
	left, right, sqrt := 0, x, 0
	for {
		if left > right {
			return sqrt
		}
		mid := left + (right - left)/2
		tmp := mid * mid
		if tmp == x {
			return mid
		} else if tmp < x {
			sqrt = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
