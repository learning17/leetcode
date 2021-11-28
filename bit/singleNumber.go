package main
// https://leetcode-cn.com/problems/WGki4K/
// 只出现一次的数字 

func singleNumber(nums []int) int {
	var arr [32]int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < 32; j++ {
			arr[j] += (nums[i] >> j) & 1
		}
	}
	var ans int32
	for i := 0; i < 32; i++ {
		if arr[i] % 3 == 1 {
			ans |= int32(1 << i)
		}
	}
	return int(ans)
}
