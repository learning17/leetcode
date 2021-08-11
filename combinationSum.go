package main

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	backtrack(&res, &path, candidates, target)
	return res
}

func backtrack(res *[][]int, path *[]int, candidates []int, target int) {
	tmpSum := sum(*path)
	if tmpSum >= target {
		if tmpSum == target {
			*res = append(*res, append([]int{}, *path...))
		}
		return
	}
	for i, candidate := range candidates {
		*path = append(*path, candidate)
		tmp := append([]int{}, candidates[i:]...)
		backtrack(res, path, tmp, target)
		*path = (*path)[:len(*path)-1]
	}
}

func sum(nums []int) int {
	sum := 0
	for _,num := range nums {
		sum += num
	}
	return sum
}
