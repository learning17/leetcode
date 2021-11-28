package main
// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
// 用最少数量的箭引爆气球

import (
	"sort"
)
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {return points[i][1] < points[j][1]})
	var tmp []int
	var num int
	for i := 0; i < len(points); i++ {
		if i == 0 || tmp[1] < points[i][0] {
			tmp = points[i]
			num++
		}
	}
	return num
}
