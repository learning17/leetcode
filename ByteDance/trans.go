package main
// https://www.nowcoder.com/practice/c3120c1c1bc44ad986259c0cf0f0b80e

import (
	"strings"
)
func trans( s string ,  n int ) string {
	arr := strings.Split(s, " ")
	var ans []string
	for i := len(arr)-1; i >= 0; i-- {
		ans = append(ans, help(arr[i]))
	}
	return strings.Join(ans, " ")
}

func help(s string) string {
	arr := []byte(s)
	for i := range arr {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			arr[i] -= 32
		} else {
			arr[i] += 32
		}
	}
	return string(arr)
}

