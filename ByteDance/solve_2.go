package main
// https://www.nowcoder.com/practice/11ae12e8c6fe48f883cad618c2e81475

import (
	"fmt"
)

func solve( s string ,  t string ) string {
	size1, size2 := len(s), len(t)
	var arr1, arr2 []byte
	if size1 > size2 {
		arr1, arr2 = []byte(s), []byte(t)
	} else {
		arr1, arr2 = []byte(t), []byte(s)
		size1, size2 = size2, size1
	}
	bit := 0
	for i := 0; i < size2; i++ {
		value := bit + int(arr1[size1-1-i]-'0' + arr2[size2-1-i]-'0')
		arr1[size1-1-i], bit = '0' + byte(value%10), value/10
	}

	for i := size1 - 1 - size2; i >= 0; i-- {
		if bit == 0 {
			break
		}
		value := bit + int(arr1[i]-'0')
		arr1[i], bit = '0' + byte(value%10), value/10
	}

	if bit != 0 {
		return fmt.Sprintf("%d%s", bit, string(arr1))
	}
	return string(arr1)
}
