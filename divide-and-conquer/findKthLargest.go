package main
// https://leetcode-cn.com/problems/xx4gT2/

func findKthLargest(nums []int, k int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	for k := size; k > 1; k-- {
		for i := (k-2)/2; i >=0; i-- {
			adjust(nums, k, i)
		}
		nums[0], nums[k-1] = nums[k-1], nums[0]
	}
	return nums[size-k]
}

func adjust(nums []int, size, i int) {
	for k := 2*i+1; k < size; k=2*k+1 {
		if k+1 < size && nums[k+1] > nums[k] {
			k++
		}
		if nums[k] > nums[i] {
			nums[i], nums[k] = nums[k], nums[i]
			i = k
		} else {
			break
		}
	}
}
