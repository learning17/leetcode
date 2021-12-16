package main
// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
/*
0～n-1中缺失的数字
*/

func missingNumber(nums []int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
