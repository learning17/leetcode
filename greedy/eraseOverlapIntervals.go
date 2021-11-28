package main
// https://leetcode-cn.com/problems/non-overlapping-intervals/
// 无重叠区间
import (
	"sort"
)
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func (i, j int) bool {return intervals[i][1] < intervals[j][1]})
	var num int
	var tmp []int
	for i := 0; i < len(intervals); i++ {
		if i == 0 || tmp[1] <= intervals[i][0] {
			tmp = intervals[i]
		} else {
			num++
		}
	}
	return num
}
