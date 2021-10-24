package main
// https://leetcode-cn.com/problems/WGki4K/
func singleNumber(nums []int) int {
	var ans int32
	for i := 0; i < 32; i++ {
		var total int32
		for _,num := range nums {
			total += (int32(num) >> i) & 1
		}
		if total % 3 > 0 {
			ans |= (1 << i)
		}
	}
	return int(ans)
}
