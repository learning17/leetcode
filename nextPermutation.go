package main

// https://leetcode-cn.com/problems/next-permutation/

func nextPermutation(nums []int)  {
	size := len(nums)
	if size <= 1 {
		return
	}
	i, j, k := size-2, size-1, size-1
	for ;i >= 0 && nums[i] >= nums[j]; i, j = i-1, j-1{
	}
	if i >= 0 {
		for ;nums[i] >= nums[k];k--{
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	for k := size-1; j < k; j,k = j+1,k-1 {
		nums[j], nums[k] = nums[k],nums[j]
	}
}
