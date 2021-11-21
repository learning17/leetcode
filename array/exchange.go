package main
// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
// 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	i, j := 0, len(nums) - 1
	for i < j {
		for ;i < j && nums[i] & 1 == 1; i++{}
		for ;i < j && nums[j] &1 == 0; j--{}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
