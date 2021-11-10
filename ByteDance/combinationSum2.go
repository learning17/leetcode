package main
// https://www.nowcoder.com/practice/75e6cd5b85ab41c6a7c43359a74e869a
import (
	"sort"
)

func combinationSum2( num []int ,  target int ) [][]int {
	sort.Ints(num)
	var res [][]int
	var path []int

	var find func([]int, int)

	find = func(arr []int, value int) {
		if value == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if value < 0 {
			return
		}
		for i, a := range arr {
			if i > 0 && arr[i] == arr[i-1] {
				continue
			}
			path = append(path, a)
			tmp := append([]int{}, arr[i+1:]...)
			find(tmp, value-a)
			path = path[:len(path)-1]
		}
	}
	find(num, target)
	return res
}

