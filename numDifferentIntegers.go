package main
// https://leetcode-cn.com/problems/number-of-different-integers-in-a-string/
import (
	"unicode"
)

func numDifferentIntegers(word string) int {
	dict := make(map[int]int)
	var value int
	for i := 0; i < len(word); {
		value, i = nextInt(word, i)
		if i == -1 {
			break
		}
		dict[value]++
	}
	return len(dict)
}

func nextInt(word string, i int) (int, int) {
	for ;i < len(word) && !unicode.IsDigit(rune(word[i]));i++ {
	}
	if i == len(word) {
		return -1, i
	}
	x := 0
	for ;i < len(word) && unicode.IsDigit(rune(word[i]));i++ {
		x = 10*x + int(word[i]-'0')
	}
	return x, i
}
