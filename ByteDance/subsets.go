package main
// https://www.nowcoder.com/practice/c333d551eb6243e0b4d92e37a06fbfc9
func subsets( A []int ) [][]int {
	var res [][]int
	var path []int

	var help func(arr []int)
	help = func(arr []int) {
		res = append(res, append([]int{}, path...))
		for i, c := range arr {
			path = append(path, c)
			tmp := append([]int{}, arr[i+1:]...)
			help(tmp)
			path = path[:len(path)-1]
		}
	}
	help(A)
	return res
}

