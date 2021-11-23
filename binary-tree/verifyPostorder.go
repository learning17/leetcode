package main
// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
// 二叉搜索树的后序遍历序列

func verifyPostorder(postorder []int) bool {
	var help func(left, right int) bool
	help = func(left, right int) bool {
		if right - left < 2 {
			return true
		}
		mid := right
		for i := right-1; i >= left; i-- {
			if postorder[i] > postorder[right] {
				mid = i
			}
			if postorder[i] > postorder[right] && postorder[i+1] < postorder[right] {
				return false
			}
		}
		return help(left, mid-1) && help(mid, right-1)
	}
	return help(0, len(postorder)-1)
}
