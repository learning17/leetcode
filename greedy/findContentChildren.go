package main
// https://leetcode-cn.com/problems/assign-cookies/
// 分发饼干
import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var num int
	for i, j := len(g)-1, len(s)-1; i >= 0 && j >= 0; {
		if s[j] >= g[i] {
			j--
			num++
		} 
		i--
	}
	return num
}
