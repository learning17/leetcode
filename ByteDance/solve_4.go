package main

// https://www.nowcoder.com/practice/fc897457408f4bbe9d3f87588f497729

import (
	"sort"
	"strconv"
)
func solve( nums []int ) string {
	help := func(a, b int) bool {
		if b == 0 {
			return true
		}
		arr_a, arr_b := []int{}, []int{}
		for ;a > 0; a /= 10 {
			arr_a = append(arr_a, a%10)
		}
		for ;b > 0; b /= 10 {
			arr_b = append(arr_b, b%10)
		}
		i, j := len(arr_a)-1, len(arr_b)-1
		for ; i >=0 && j >= 0; i, j = i-1, j-1 {
			if arr_a[i] > arr_b[j] {
				return true
			}
		}
		if i >= 0 && arr_a[i] > arr_b[len(arr_b)-1] {
			return true
		}
		return false
	}
	sort.Slice(nums, func(i, j int) bool {return help(nums[i], nums[j])})
	if len(nums) == 0 {
		return ""
	}
	ans := ""
	for i := 0; i < len(nums); i++ {
		if len(ans) == 0 {
			if nums[i] != 0 {
				strconv.Itoa(nums[i])
			}
			continue
		}
		ans += strconv.Itoa(nums[i])
	}
	return ans
}

