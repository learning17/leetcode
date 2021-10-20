package main

// https://www.nowcoder.com/practice/fc897457408f4bbe9d3f87588f497729

import (
	"sort"
	"strconv"
	"fmt"
)
func main() {
	nums := []int{10,1}
	res := solve(nums)
	fmt.Println(res)
}

func solve( nums []int ) string {
	help := func(a, b int) bool {
		if a == 0 {
			return false
		}
		if b == 0 {
			return true
		}
		arr_a, arr_b := []int{}, []int{}
		for ; a > 0; a /= 10 {
			arr_a = append(arr_a, a % 10)
		}
		for ; b > 0; b /= 10 {
			arr_b = append(arr_b, b % 10)
		}
		i, j := len(arr_a)-1, len(arr_b)-1
		for ;i >=0 && j >= 0; i, j = i-1, j-1 {
			if arr_a[i] > arr_b[j] {
				return true
			}
			if arr_a[i] < arr_b[j] {
				return false
			}
		}
		if i >= 0 && arr_a[i] > arr_b[0] {
			return true
		}
		if j >= 0 && arr_b[j] < arr_a[0] {
			return true
		}
		return false
	}
	sort.Slice(nums, func(i, j int) bool {return help(nums[i], nums[j])})
	if nums[0] == 0 {
		return "0"
	}
	ans := ""
	for i, num := range nums {
		if i == 0 {
			ans = strconv.Itoa(num)
		} else {
			ans += strconv.Itoa(num)
		}
	}
	return ans
}

