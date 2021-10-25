package main
// https://www.nowcoder.com/practice/c9addb265cdf4cdd92c092c655d164ca
import (
	"strings"
)

func generateParenthesis( n int ) []string {
	var res, path []string
	var help func(left int, right int) 
	help = func(left int, right int) {
		if left > right || left < 0 || right < 0 {
			return
		}
		if left == 0 && right == 0 {
			res = append(res, strings.Join(path,""))
			return
		}
		path = append(path, "(")
		help(left-1, right)
		path = path[:len(path)-1]
		path = append(path, ")")
		help(left, right-1)
		path = path[:len(path)-1]
	}
	help(n, n)
	return res
}

