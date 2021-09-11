package main
import "strconv"
// https://leetcode-cn.com/problems/different-ways-to-add-parentheses/

func diffWaysToCompute(expression string) []int {
	if intV, err := strconv.Atoi(expression); err == nil {
		return []int{intV}
	}
	res := []int{}
	for i, e := range expression {
		if e == '+' || e == '-' || e == '*' {
			left  := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _,l := range left {
				for _, r := range right {
					tmpRes := 0
					if e == '+' {
						tmpRes = l + r
					} else if e == '-' {
						tmpRes = l - r
					} else {
						tmpRes = l * r
					}
					res = append(res, tmpRes)
				}
			}
		}
	}
	return res
}
