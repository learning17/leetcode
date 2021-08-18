package main
// https://leetcode-cn.com/problems/valid-parentheses/

func isValid(s string) bool {
	size := len(s)
	if size == 0 {
		return false
	}
	stack := make([]byte, 0, size)
	for i := 0; i < size; i++ {
		count := len(stack)
		if count == 0 || s[i] == '(' || s[i] == '[' || s[i] == '{'{
			stack = append(stack, s[i])
			continue
		}
		if s[i] == ')' && stack[count-1] != '(' {
			return false
		}
		if s[i] == ']' && stack[count-1] != '[' {
			return false
		}
		if s[i] == '}' && stack[count-1] != '{' {
			return false
		}
		stack = stack[:count-1]
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
