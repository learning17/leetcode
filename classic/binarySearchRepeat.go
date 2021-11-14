package main
// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// 最左侧和最右侧元素二分查找

func searchRange(nums []int, target int) []int {
	return []int{searchLeft(nums, target), searchRight(nums, target)}
}

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] >= target {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left+(right-left)/2
		if nums[mid] <= target {
			left = mid+1
		} else {
			right = mid-1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
