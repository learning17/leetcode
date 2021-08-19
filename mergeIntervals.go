package main

// https://leetcode-cn.com/problems/merge-intervals/

func merge(intervals [][]int) [][]int {
	size := len(intervals)
	if size <= 1 {
		return intervals
	}
	quickSort(intervals, 0 , size-1)
	res := [][]int{intervals[0]}
	j := 0
	for i := 1; i < size; i++ {
		if res[j][1] < intervals[i][0] {
			res = append(res, intervals[i])
			j++
		} else if res[j][1] >= intervals[i][0] && res[j][1] <= intervals[i][1] {
			res[j][1] = intervals[i][1]
		}
	}
	return res
}

func quickSort(nums [][]int, left, right int){
	deal(nums, left, right)
	if right - left < 3 {
		return
	}
	i, j := left+1, right-2
	for{
		for ;nums[i][0] < nums[right-1][0];i++ {
		}
		for ;nums[j][0] >= nums[right-1][0] && j > i; j--{
		}
		if j <= i {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[right-1] = nums[right-1], nums[i]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func deal(nums [][]int, left, right int) {
	if left >= right {
		return
	}
	mid := left +(right - left)/2
	if nums[left][0] > nums[mid][0] {
		nums[left], nums[mid] = nums[mid], nums[left]
	}
	if nums[mid][0] > nums[right][0] {
		nums[mid], nums[right] = nums[right], nums[mid]
	}
	if nums[left][0] > nums[mid][0] {
		nums[left], nums[mid] = nums[mid], nums[left]
	}
	nums[mid], nums[right-1] = nums[right-1], nums[mid]
}
