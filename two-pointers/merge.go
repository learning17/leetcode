package main
// https://leetcode-cn.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j, k := m-1, n-1, m+n-1
	for {
		if i < 0 || j < 0 {
			break
		}
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	if j >= 0 {
		copy(nums1[:j+1], nums2[:j+1])
	}
}
