package main
import "strings"
// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var path []string
	var res []string

	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		if right < left || left < 0 || right < 0 {
			return
		}
		if left == 0 && right == 0 {
			res = append(res, strings.Join(path,""))
			return
		}
		path = append(path, "(")
		backtrack(left-1, right)
		path = path[:len(path)-1]
		path = append(path, ")")
		backtrack(left, right-1)
		path = path[:len(path)-1]
	}
	backtrack(n, n)
	return res
}
