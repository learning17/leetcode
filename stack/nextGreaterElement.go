package main
// https://leetcode-cn.com/problems/next-greater-element-i/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	dict := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		if len(stack) == 0 {
			stack = append(stack, nums2[i])
			continue
		}
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			dict[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for i := 0; i < len(stack); i++ {
		dict[stack[i]] = -1
	}
	var res []int
	for _,num := range nums1 {
		res = append(res, dict[num])
	}
	return res
}
