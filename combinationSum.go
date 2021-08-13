package main

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var backtrack func([]int, int)

	backtrack = func(nums []int, subTarget int) {
		if subTarget <= 0 {
			if subTarget == 0 {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		for i, num := range nums {
			path = append(path, num)
			tmp := append([]int{}, nums[i:]...)
			backtrack(tmp, subTarget - num)
			path = path[:len(path)-1]
		}
	}
	backtrack(candidates, target)
	return res
}

