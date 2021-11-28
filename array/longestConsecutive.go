package main
// https://leetcode-cn.com/problems/WhsWhI/
// 最长连续序列

func longestConsecutive(nums []int) int {
	dict := make(map[int]bool)
	for _, num := range nums {
		dict[num] = true
	}
	visited := make(map[int]bool)
	var maxSize int
	for _, num := range nums {
		if _, ok := visited[num]; ok {
			continue
		}
		var size int
		for i := 0;; i++ {
			tmp := num+i
			if _, ok := dict[tmp]; !ok {
				break
			}
			visited[tmp] = true
			size++
		}
		if size > maxSize {
			maxSize = size
		}
	}
	return maxSize
}
