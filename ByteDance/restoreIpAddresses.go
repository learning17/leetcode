package main
// https://www.nowcoder.com/practice/ce73540d47374dbe85b3125f57727e1e

import (
	"strings"
	"strconv"
	"fmt"
)

func main() {
	s := "010010"
	res := restoreIpAddresses(s)
	fmt.Println(res)
}
func restoreIpAddresses( s string ) []string {
	var res, path []string
	var help func(n int, str string)
	help = func(n int, str string) {
		if n == 0 && len(str) == 0 {
			fmt.Println(path)
			res = append(res, strings.Join(path, "."))
			return
		}
		if n == 0 || len(str) == 0 {
			return
		}
		if str[0] == '0' {
			path = append(path, "0")
			help(n-1, str[1:])
			path = path[:len(path)-1]
			return
		}
		for i := 0; i < len(str); i++ {
			if v, err := strconv.Atoi(str[:i+1]); err != nil || v > 255 {
				break
			}
			path = append(path, str[:i+1])
			help(n-1, str[i+1:])
			path = path[:len(path)-1]
		}
	}
	help(4, s)
	return res
}

