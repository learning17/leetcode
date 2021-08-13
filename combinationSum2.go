package main

import "sort"
// https://leetcode-cn.com/problems/combination-sum-ii/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
    var res [][]int
    var path []int
    var backtrack func([]int, int)
    backtrack = func (nums []int, subTarget int) {
        if subTarget <= 0 {
            if subTarget == 0 {
                res = append(res, append([]int{}, path...))
            }
            return
        }
        for i, num := range nums {
			if i != 0 && nums[i] == nums[i-1] {
				continue
			}
            path = append(path, num)
			tmp := append([]int{}, nums[i+1:]...)
            backtrack(tmp, subTarget - num)
            path = path[:len(path)-1]
        }
    }
    backtrack(candidates, target)
    return res
}

