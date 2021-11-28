package main
// https://leetcode-cn.com/problems/partition-labels/
// 划分字母区间

func partitionLabels(s string) []int {
	dict := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dict[s[i]] = i
	}
	var start, end int
	var res []int
	for i := 0; i < len(s); i++ {
		if dict[s[i]] > end {
			end = dict[s[i]]
		}
		if i == end {
			res = append(res, end-start+1)
			start, end = i+1, i+1
		}
	}
	return res
}
