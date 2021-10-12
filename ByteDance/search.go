package main

// https://www.nowcoder.com/practice/4f470d1d3b734f8aaf2afb014185b395

func search( nums []int ,  target int ) int {
	var help func(int, int) int
	help = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := left + (right-left)/2
		if nums[mid] >= target {
			if left == mid {
				if nums[mid] == target {
					return left
				} else {
					return -1
				}
			}
			return help(left, mid)
		} else {
			return help(mid+1, right)
		}
	}
	return help(0, len(nums)-1)
}

