package main
// https://www.nowcoder.com/practice/44d8c152c38f43a1b10e168018dcc13f


func atoi( str string ) int {
	if len(str) == 0 {
		return 0
	}
	idx, sign, ans := 0, 1, 0
	for ;str[idx] == ' ';idx++{
	}
	if str[idx] == '-' {
		sign = -1
		idx++
	}
	if str[idx] == '+' {
		idx++
	}
	for ;idx < len(str);idx++ {
		if str[idx] < '0' || str[idx] > '9' {
			break
		}
		ans = 10*ans + int(str[idx] - '0')
	}
	return ans*sign
}

