package main
// 颜色分类
// https://leetcode.cn/problems/sort-colors/
/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
必须在不使用库的sort函数的情况下解决这个问题。
*/

func sortColors(nums []int)  {
    left, right := 0, len(nums)-1
    for i := 0; i <= right; i++ {
        for ;i <= right && nums[i] == 2; right-- {
            nums[i], nums[right] = nums[right], nums[i]
        }
        if nums[i] == 0 {
            nums[i], nums[left] = nums[left], nums[i]
            left++
        }
    }
}


