package main

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/submissions/
func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums) - 1)
	return nums[len(nums)-k]
}

func quickSort(nums []int, left, right int) {
	dealPivot(nums, left, right)
	if right - left < 3 {
		return
	}
	i, j := left+1, right-2
	for {
		for ; nums[i] < nums[right-1]; i++ {
		} 
		for ; nums[j] >= nums[right-1] && j > i; j-- {
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[right-1] = nums[right-1], nums[i]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func dealPivot(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right - left)/2
	if nums[left] > nums[mid] {
		nums[left], nums[mid] = nums[mid], nums[left]
	}
	if nums[mid] > nums[right] {
		nums[mid], nums[right] = nums[right], nums[mid]
	}
	if nums[left] > nums[mid] {
		nums[left], nums[mid] = nums[mid], nums[left]
	}
	nums[mid], nums[right-1] = nums[right-1], nums[mid]
}

