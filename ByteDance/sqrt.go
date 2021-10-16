package main
// https://www.nowcoder.com/practice/09fbfb16140b40499951f55113f2166c

func sqrt( x int ) int {
	left, right := 0, x
	for {
		if left > right {
			return right
		}
		mid := left + (right-left)/2
		tmp := mid*mid
		if tmp == x {
			return mid
		} else if tmp < x {
			left = mid+1
		} else {
			right = mid-1
		}
	}
}
