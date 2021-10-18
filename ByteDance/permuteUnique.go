package main
// https://www.nowcoder.com/practice/a43a2b986ef34843ac4fdd9159b69863

import "sort"
func permuteUnique( num []int ) [][]int {
	sort.Ints(num)
	res := [][]int{}
	path := []int{}
	size := len(num)
	var help func([]int)
	help = func(arr []int) {
		if len(path) == size {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := range arr {
			if i > 0 && arr[i] == arr[i-1] {
				continue
			}
			path = append(path, arr[i])
			tmp := append([]int{}, arr[:i]...)
			tmp = append(tmp, arr[i+1:]...)
			help(tmp)
			path = path[:len(path)-1]
		}
	}
	help(num)
	return res
}
