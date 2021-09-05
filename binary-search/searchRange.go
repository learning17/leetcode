package main
// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for {
		if left > right {
			return -1
		}
		if left == right {
			if nums[left] == target {
				return left
			} else {
				return -1
			}
		}
		mid := left + (right - left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
}

func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		if left > right {
			return -1
		}
		if left == right {
			if nums[right] == target {
				return right
			} else {
				return -1
			}
		}
		mid := left + (right - left+1)/2
		if nums[mid] <= target {
			left = mid
		} else{
			right = mid-1
		}
	}
}
