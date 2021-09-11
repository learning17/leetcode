package main
// https://leetcode-cn.com/problems/permutation-i-lcci/

func permutation(S string) []string {
	arr := []byte(S)
	var path []byte
	var res []string

	var backtrack func([]byte)
	size := len(arr)
	backtrack = func(subArr []byte) {
		if len(path) == size {
			res = append(res, string(path))
			return
		}
		for i, a := range subArr {
			path = append(path, a)
			tmp := append([]byte{}, subArr[:i]...)
			tmp = append(tmp, subArr[i+1:]...)
			backtrack(tmp)
			path = path[:len(path)-1]
		}
	}
	backtrack(arr)
	return res
}
