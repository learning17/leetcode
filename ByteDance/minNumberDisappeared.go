package main
// https://www.nowcoder.com/practice/50ec6a5b0e4e45348544348278cdcee5

func minNumberDisappeared( nums []int ) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] <= 0 {
			nums[i] = size+1
		}
	}
	for i := 0; i < size; i++ {
		if abs(nums[i]) <= size {
			nums[abs(nums[i])-1] = -abs(nums[abs(nums[i])-1])
		}
	}
	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			return i+1
		}
	}
	return size+1
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

