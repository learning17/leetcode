package main
// https://www.nowcoder.com/practice/fe6b651b66ae47d7acce78ffdd9a96c7

import (
	"sort"
)

func Permutation( str string ) []string {
	res := []string{}
	path := []byte{}
	size := len(str)

	var help func([]byte)
	help = func(arr []byte) {
		if len(path) == size {
			res = append(res, string(path))
			return
		}
		for i,c := range arr {
			if i > 0 && arr[i] == arr[i-1] {
				continue
			}
			path = append(path, c)
			tmp := append([]byte{}, arr[:i]...)
			tmp = append(tmp, arr[i+1:]...)
			help(tmp)
			path = path[:len(path)-1]
		}
	}
	arr := []byte(str)
	sort.Slice(arr, func(i, j int) bool {return arr[i] < arr[j]})
	help(arr)
	return res
}

