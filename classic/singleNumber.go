package main

// https://leetcode-cn.com/problems/WGki4K/
// 只出现一次的数字

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		var tmp int
		for _, num := range nums {
			tmp += (num >> i) & 1
		}
		if tmp%3 == 1 {
			ans += (1 << i)
		}
	}
	return int(ans)
}
