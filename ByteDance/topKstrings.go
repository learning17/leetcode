package main
// https://www.nowcoder.com/practice/fd711bdfa0e840b381d7e1b82183b3ee

import (
	"strconv"
	"fmt"
)

func topKstrings( strings []string ,  k int ) [][]string {
	dict := make(map[string]int)
	for _, s := range strings {
		dict[s]++
	}
	arr := make([][]string, 0, len(dict))
	for s, v := range dict {
		arr = append(arr, []string{s, strconv.Itoa(v)})
	}
	for j := (k-2)/2; j >= 0; j-- {
		adjust(arr, j, k)
	}

	for j := k; j < len(arr); j++ {
		if cmp(arr[j], arr[0]) {
			continue
		}
		arr[0] = arr[j]
		adjust(arr, 0, k)
	}
	ans := make([][]string, k)
	for i := k-1; i >= 0; i-- {
		ans[i] = arr[0]
		arr[0] = arr[i]
		adjust(arr, 0, i)
	}
	return ans
}

func cmp(node1, node2 []string) bool {
	intV1, _ := strconv.Atoi(node1[1])
	intV2, _ := strconv.Atoi(node2[1])
	if intV1 < intV2 {
		return true
	}
	if intV1 > intV2 {
		return false
	}
	return node1[0] > node2[0]
}
func adjust(arr [][]string, k, size int) {
	for i := 2*k+1; i < size; k, i = i, 2*i+1 {
		if i + 1 < size && !cmp(arr[i], arr[i+1]) {
			i++
		}
		if cmp(arr[k], arr[i]) {
			break
		}
		arr[k], arr[i] = arr[i], arr[k]
	}
}
