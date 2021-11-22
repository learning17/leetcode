package main
// https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
// 把数组排成最小的数
import (
	"sort"
	"bytes"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		num1, num2 := 10*nums[i], 10*nums[j]
		for num := nums[i] / 10; num > 0; num /= 10 {
			num2 *= 10
		}
		for num := nums[j] / 10; num > 0; num /= 10 {
			num1 *= 10
		}
		num1, num2 = num1+nums[j], num2+nums[i]
		return num1 < num2
	})
	var buf bytes.Buffer
	for i := 0; i < len(nums); i++ {
		buf.WriteString(strconv.Itoa(nums[i]))
	}
	return buf.String()
}
