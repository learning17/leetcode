package main

// https://leetcode-cn.com/problems/subsets/

func subsets(nums []int) [][]int {
	var path []int
	var res [][]int

	var backtrack func([]int)
	backtrack = func(arr []int) {
		res = append(res, append([]int{}, path...))

		for i, a := range arr {
			path = append(path, a)
			tmp := append([]int{}, arr[i+1:]...)
			backtrack(tmp)
			path = path[:len(path)-1]
		}
	}
	backtrack(nums)
	return res
}

