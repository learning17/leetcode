package main
import (
	"strconv"
	"unicode"
)
// https://leetcode-cn.com/problems/different-ways-to-add-parentheses/

func diffWaysToCompute(expression string) []int {
	if v, err := strconv.Atoi(expression); err == nil {
		return []int{v}
	}
	var res []int
	for i := 0; i < len(expression); i++ {
		if unicode.IsDigit(rune(expression[i])) {
			continue
		}
		left := diffWaysToCompute(expression[:i])
		right := diffWaysToCompute(expression[i+1:])
		for _,l := range left {
			for _,r := range right {
				switch expression[i] {
				case '+':
					res = append(res, l+r)
				case '-':
					res = append(res, l-r)
				case '*':
					res = append(res, l*r)
				}
			}
		}
	}
	return res
}
