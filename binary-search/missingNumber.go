package main
// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

func missingNumber(nums []int) int {
	left, right := 0, len(nums)
	for {
		if left == right {
			return left
		}
		mid := left + (right - left)/2
		if nums[mid] == mid {
			left = mid+1
		} else {
			right = mid
		}
	}
}
