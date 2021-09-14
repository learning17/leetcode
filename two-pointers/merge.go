package main

// https://leetcode-cn.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m+n-1
	i, j := m-1, n-1
	for ; i >= 0 && j >= 0; {
		if nums1[i] > nums2[j] {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}
	if j >= 0 {
		copy(nums1[:j+1], nums2[0:j+1])
	}
}
