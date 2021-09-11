package main
// https://leetcode-cn.com/problems/combinations/

func combine(n int, k int) [][]int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i+1
	}
	var path []int
	var res [][]int

	var backtrack func([]int, int)
	backtrack = func(arr []int, k int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i, a := range arr {
			path = append(path, a)
			tmp := append([]int{}, arr[i+1:]...)
			backtrack(tmp, k)
			path = path[:len(path)-1]
		}
	}
	backtrack(nums, k)
	return res
}
