package main
// https://leetcode-cn.com/problems/xx4gT2/

func findKthLargest(nums []int, k int) int {
	return 0
}

// 堆排序
func heapSort(nums []int) {
	for size := len(nums); size > 1; size-- {
		for k := (size-2)/2; k >= 0; k-- {
			adjust(nums, size, k)
		}
		nums[0], nums[size-1] = nums[size-1], nums[0]
	}
}

func adjust(nums []int, size int, k int) {
	for i := 2*k+1; i < size; i, k = 2*i+1, i {
		if i+1 < size && nums[i] < nums[i+1] {
			i++
		}
		if nums[i] < nums[k] {
			break
		}
		nums[k], nums[i] = nums[i], nums[k]
	}
}

// 快排
func quickSort(nums []int, left int, right int) {
	help(nums, left, right)
	if right - left < 3 {
		return
	}
	i, j := left+1, right-2
	for {
		for ;nums[i] < nums[right-1];i++{
		}
		for ;i < j && nums[j] >= nums[right-1];j--{
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[right-1] = nums[right-1],nums[i]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func help(nums []int, left int, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
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

