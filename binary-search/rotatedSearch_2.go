package main
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/
/*
搜索旋转排序数组 II
数组中的值不必互不相同
*/
func search(nums []int, target int) bool {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}