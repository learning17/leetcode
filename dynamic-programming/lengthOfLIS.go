package main
// https://leetcode-cn.com/problems/longest-increasing-subsequence/

import "math"

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	ends := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		ends[i] = math.MaxInt64
	}
	ends[0], ends[1] = 0, nums[0]
	maxSize := 1
	for i := 1; i < size; i++ {
		pos := findBinary(ends, nums[i])
		if ends[pos+1] > nums[i] {
			ends[pos+1] = nums[i]
		}
		if pos+1 > maxSize {
			maxSize = pos+1
		}
	}
	return maxSize
}

func findBinary(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		if left >= right {
			return left
		}
		mid := left + (right-left+1)/2
		if nums[mid] >= target {
			right = mid-1
		} else {
			left = mid
		}
	}
}


/*func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp := make([]int, size)
	maxSize := 1
	for i := 0; i < size; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxSize = max(maxSize, dp[i])
	}
	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}*/
