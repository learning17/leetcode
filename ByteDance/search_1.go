package main
// https://www.nowcoder.com/practice/87c0e7abcbda41e7963660fa7d020995

func search( nums []int ,  target int ) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} 
		if nums[mid] >= nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid-1
			} else {
				left = mid+1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid+1
			} else {
				right = mid -1
			}
		}
	}
	return -1
}

