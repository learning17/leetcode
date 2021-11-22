package main
// https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
// 扑克牌中的顺子
import (
	"math"
)
func isStraight(nums []int) bool {
	max, min := nums[0], math.MaxInt64 
	dict := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if _, ok := dict[nums[i]]; ok {
			return false
		}
		dict[nums[i]] = true
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return max - min < 5
}
