package main
// https://leetcode-cn.com/problems/candy/
// 分发糖果

func candy(ratings []int) int {
	size := len(ratings)
	if size < 2 {
		return size
	}
	candys := make([]int, size)
	for i := 0; i < size; i++ {
		candys[i] = 1
	}
	for i := 1; i < size; i++ {
		if ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1]+1
		}
	}
	for i := size - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candys[i] <= candys[i+1] {
			candys[i] = candys[i+1] + 1
		}
	}
	var sum int
	for i := 0; i < size; i++ {
		sum += candys[i]
	}
	return sum
}
