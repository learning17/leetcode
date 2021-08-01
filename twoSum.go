package main

// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	d := make(map[int]int)
	for i, num := range nums {
		if pos, ok := d[target-num]; ok {
			return []int{pos, i}
		}
		d[num] = i
	}
	return []int{} 
}
