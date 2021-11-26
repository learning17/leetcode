package main
// https://leetcode-cn.com/problems/8Zf90G/
// 后缀表达式
import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for i := 0; i < len(tokens); i++ {
		if v, err := strconv.Atoi(tokens[i]); err == nil {
			stack = append(stack, v)
			continue
		}
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		switch tokens[i] {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, a-b)
		case "*":
			stack = append(stack, a*b)
		case "/":
			stack = append(stack, a/b)
		}
	}
	return stack[0]
}
