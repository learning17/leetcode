package main
// https://leetcode-cn.com/problems/3sum/
// 三数之和
import (
	"sort"
)
func threeSum(nums []int) [][]int {
	var res [][]int
	size := len(nums)
	if size < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < size; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, size - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == -nums[i] {
				tmp := []int{nums[i], nums[left], nums[right]}
				res = append(res, tmp)
				for ;left < size && nums[left] == tmp[1];left++{}
				for ;right >= 0 && nums[right] == tmp[2];right--{};
			} else if sum > -nums[i] {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
