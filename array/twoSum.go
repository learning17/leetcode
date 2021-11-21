package main
// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
// 和为s的两个数字
func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] + nums[j] == target {
			break
		} else if nums[i] + nums[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{nums[i], nums[j]}
}
