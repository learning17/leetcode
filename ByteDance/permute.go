package main
// https://www.nowcoder.com/practice/4bcf3081067a4d028f95acee3ddcd2b1
func permute( num []int ) [][]int {
	var res [][]int
	var path []int
	size := len(num)

	var help func([]int)
	help = func(arr []int) {
		if len(path) == size {
			res = append(res, append([]int{}, path...))
			return
		}
		for i, c := range arr {
			path = append(path, c)
			tmp := append([]int{}, arr[:i]...)
			tmp = append(tmp, arr[i+1:]...)
			help(tmp)
			path = path[:len(path)-1]
		}
	}
	help(num)
	return res
}
