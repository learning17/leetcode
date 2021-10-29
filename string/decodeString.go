package main
//https://leetcode-cn.com/problems/decode-string/
import (
	"strconv"
	"unicode"
	"strings"
)

func decodeString(s string) string {
	var stackS []string
	var stackN []int
	size := len(s)
	for i := 0; i < size; {
		if s[i] == '[' || unicode.IsLetter(rune(s[i])){
			stackS = append(stackS, s[i:i+1])
			i++
		} else if unicode.IsDigit(rune(s[i])) {
			pos := i
			for ; i < size && unicode.IsDigit(rune(s[i]));i++{
			}
			v,_ := strconv.Atoi(s[pos:i])
			stackN = append(stackN, v)
		} else {
			pos := len(stackS) - 1
			for ;pos >= 0 && stackS[pos] != "["; pos-- {
			}
			newStr := strings.Join(stackS[pos+1:], "")
			stackS = stackS[:pos]
			stackS = append(stackS, strings.Repeat(newStr, stackN[len(stackN)-1]))
			stackN = stackN[:len(stackN)-1]
			i++
		} 
	}
	return strings.Join(stackS, "")
}
