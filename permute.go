package main
// https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	backtrack(&res, &path, nums, len(nums))
	return res
}

func backtrack(res *[][]int, path *[]int, nums []int, size int) {
	if len(*path) == size {
		tmp := make([]int, size)
		copy(tmp, *path)
		*res = append(*res, tmp)
		return
	}
	for i,num := range nums {
		*path = append(*path, num)
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		tmp = append(tmp[:i], tmp[i+1:]...)
		backtrack(res, path, tmp, size)
		*path = (*path)[:len(*path)-1]
	}
}
