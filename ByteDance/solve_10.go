package main
// https://www.nowcoder.com/practice/c4c488d4d40d4c4e9824c3650f7d5571

import (
	"strconv"
)

func solve( s string ,  t string ) string {
	var arrS, arrT []int
	for i := len(s); i > 0; i-- {
		v, _ := strconv.Atoi(s[i-1:i])
		arrS = append(arrS, v)
	}
	for i := len(t); i > 0; i-- {
		v, _ := strconv.Atoi(t[i-1:i])
		arrT = append(arrT, v)
	}
	arr := make([]int, len(arrS)+len(arrT))
	for i := 0; i < len(arrS); i++ {
		for j := 0; j < len(arrT); j++ {
			arr[i+j] += arrS[i]*arrT[j]
		}
	}
	size := len(arr)
	for i := 0; i < size-1; i++ {
		arr[i+1] += arr[i] / 10
		arr[i] = arr[i] % 10
	}
	var ans string
	for i := size-1; i >= 0; i-- {
		if i == size - 1 && arr[i] == 0 {
			continue
		}
		ans += strconv.Itoa(arr[i])
	}
	return ans
}

