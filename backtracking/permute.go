package main
// https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
	var path []int
	var res [][]int
	var backtrack func([]int, int)

	backtrack = func(arr []int, size int) {
		if len(path) == size {
			res = append(res, append([]int{}, path...))
			return
		}
		for i, a := range arr {
			path = append(path, a)
			tmp := append([]int{}, arr[:i]...)
			tmp = append(tmp, arr[i+1:]...)
			backtrack(tmp, size)
			path = path[:len(path)-1]
		}
	}
	backtrack(nums, len(nums))
	return res
}

