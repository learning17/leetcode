package main
// https://www.nowcoder.com/practice/45fd68024a4c4e97a8d6c45fc61dc6ad
import "fmt"
func main() {
	res := longestValidParentheses(")()(()()((((((())(")
	fmt.Println(res)
}
func longestValidParentheses( s string ) int {
	stack := []int{-1}
	var maxSize int
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if maxSize < i - stack[len(stack)-1] {
					maxSize = i - stack[len(stack)-1]
				}
			}
		}
	}
	return maxSize
}

