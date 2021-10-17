package main
// https://www.nowcoder.com/practice/fd711bdfa0e840b381d7e1b82183b3ee

import (
	"sort"
	"strconv"
)

func topKstrings( strings []string ,  k int ) [][]string {
	dict := make(map[string]int)
	for _, s := range strings {
		dict[s]++
	}
	type node struct {
		key string
		value int
	}
	arr := make([]node, 0, len(dict))
	for k, v := range dict {
		arr = append(arr, node{k,v})
	}
	help := func(n1, n2 node) bool {
		if n1.value > n2.value {
			return true
		}
		if n1.value == n2.value && n1.key < n2.key {
			return true
		}
		return false
	}
	sort.Slice(arr, func(i,j int) bool {return help(arr[i], arr[j])})
	ans := [][]string{}
	for i := 0; i < k; i++ {
		ans = append(ans, []string{arr[i].key, strconv.Itoa(arr[i].value)})
	}
	return ans
}
