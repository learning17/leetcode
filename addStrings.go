package main

import "strconv"
// https://leetcode-cn.com/problems/add-strings/

func addStrings(num1 string, num2 string) string {
	add := 0
	res := ""
	for i,j := len(num1)-1,len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i]-'0')
		}
		if j >= 0 {
			y = int(num2[j]-'0')
		}
		tmp := add + x + y
		res = strconv.Itoa(tmp % 10) + res
		add = tmp/10
	}
	return res
}
