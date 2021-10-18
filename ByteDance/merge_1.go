package main
// https://www.nowcoder.com/practice/69f4e5b7ad284a478777cb2a17fb5e6a

import (
	"sort"
)

type Interval struct {
	Start int
	End int
}

func merge( intervals []*Interval ) []*Interval {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {return intervals[i].Start < intervals[j].Start})
	ans := []*Interval{}
	node := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if node.End >= intervals[i].Start && node.End <= intervals[i].End {
			node.End = intervals[i].End
		} else if intervals[i].Start > node.End {
			ans = append(ans, node)
			node = intervals[i]
		}
	}
	ans = append(ans, node)
	return ans
}

