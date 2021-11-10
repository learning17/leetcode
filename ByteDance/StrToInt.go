package main
// https://www.nowcoder.com/practice/d11471c3bf2d40f38b66bb12785df47f

func StrToInt( s string ) int {
	size := len(s)
	if size == 0 {
		return 0
	}
	var i int
	for ;i < size && s[i] ==' '; i++{
	}
	if i == size || (s[i] != '-' && s[i] != '+' && (s[i] < '0' || s[i] > '9')) {
		return 0
	} 
	sign := 1
	if s[i] == '-' {
		sign = -1
	}
	if s[i] == '-' || s[i] == '+' {
		i++
	}
	var ans int
	maxValue := 1 << 31
	for ;i < size; i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		ans = 10*ans + int(s[i] - '0')
		if ans >= maxValue {
			ans = maxValue
			break
		}
	}
	ans *= sign
	if ans == maxValue {
		ans--
	}
	return ans
}

