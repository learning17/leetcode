package main
// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
/*
在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。
找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
*/

func searchRange(nums []int, target int) []int {
	return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
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

func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
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
