package main
// https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
// 在排序数组中查找数字 I

func search(nums []int, target int) int {
	left := searchLeft(nums, target)
	right := searchRight(nums, target)
	if left == -1 || right == -1 {
		return 0
	}
	return searchRight(nums, target) - searchLeft(nums, target) + 1
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
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}


