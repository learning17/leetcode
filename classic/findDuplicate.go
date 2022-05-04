package main

// https://leetcode-cn.com/problems/find-the-duplicate-number/submissions/
// 寻找重复数

func findDuplicate(nums []int) int {
	var slow, fast int
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}
