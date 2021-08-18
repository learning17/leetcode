package main
// https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	size := len(nums)
	var backtrack func([]int)
	backtrack = func(subNums []int) {
		if len(path) == size {
			res = append(res, append([]int{}, path...))
			return
		}
		for i, num := range subNums {
			path = append(path, num)
			tmp := append([]int{}, subNums...)
			tmp = append(tmp[:i], tmp[i+1:]...)
			backtrack(tmp)
			path = path[:len(path)-1]
		}
	}
	backtrack(nums)
	return res
}

