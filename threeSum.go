package main

import (
	"sort"
	"fmt"
)
// https://leetcode-cn.com/problems/3sum/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	res := [][]int{}
	d := make(map[string]bool)
	for i := 0; i < size; i++ {
		left, right := i+1, size-1
		for ; right > left; {
			tmp := nums[left] + nums[right]
			if tmp == -nums[i] {
				key := fmt.Sprintf("%d_%d_%d", nums[i], nums[left], nums[right])
				if _,ok := d[key]; !ok {
					res = append(res, []int{nums[i], nums[left], nums[right]})
					d[key] = true
				}
				left++
				right--
			} else if tmp < -nums[i] {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
