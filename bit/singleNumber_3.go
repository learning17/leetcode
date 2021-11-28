package main
// https://leetcode-cn.com/problems/single-number-iii/
// 只出现一次的数字 III

func singleNumber(nums []int) []int {
	var tmp int
	for i := 0; i < len(nums); i++ {
		tmp ^= nums[i]
	}
	var bit int
	for ;tmp != 0 && tmp & 1 == 0; tmp >>= 1 {
		bit++
	}
	ans := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		if (nums[i] >> bit) & 1 == 1 {
			ans[0] ^= nums[i]
		} else {
			ans[1] ^= nums[i]
		}
	}
	return ans
}
