package main

// https://www.nowcoder.com/practice/37548e94a270412c8b9fb85643c8ccc2

func isValid( s string ) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			size := len(stack)
			if size == 0 {
				return false
			}
			tmpC := stack[size-1]
			stack = stack[:size-1]
			if !cmp(tmpC, c) {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func cmp (c1, c2 rune) bool {
	if c1 == '(' && c2 == ')' {
		return true
	}
	if c1 == '[' && c2 == ']' {
		return true
	}
	if c1 == '{' && c2 == '}' {
		return true
	}
	return false
}

