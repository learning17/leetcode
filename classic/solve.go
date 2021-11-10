package main
// https://www.nowcoder.com/practice/2cc32b88fff94d7e8fd458b8c7b25ec1
// 进制转换

import (
	"bytes"
)


func solve( M int ,  N int ) string {
	var ans bytes.Buffer
	if M < 0 {
		ans.WriteByte('-')
		M = -M
	}
	var arr []int
	for ;M > 0; M /= N {
		arr = append(arr, M % N)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= 10 {
			ans.WriteByte('A' + byte(arr[i]-10))
		} else {
			ans.WriteByte('0' + byte(arr[i]))
		}
	}
	return ans.String()
}

