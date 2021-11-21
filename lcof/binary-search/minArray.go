package main
// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
// 旋转数组的最小数字

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + (right-left)/2
		if numbers[mid] > numbers[right] {
			left = mid+1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right--
		}
	}
	return numbers[left]
}
