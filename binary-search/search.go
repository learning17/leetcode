package main

// https://leetcode-cn.com/problems/binary-search/

func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for {
		if left > right {
			return -1
		}
		mid := left + (right - left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid+1
		}
	}
}

