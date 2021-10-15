package main
// https://www.nowcoder.com/practice/c3a6afee325e472386a1c4eb1ef987f3

func solve( str string ) string {
	size := len(str)
	if size == 0 {
		return ""
	}
	res := make([]byte, size)
	for i := 0; i < size; i++ {
		res[size-i-1] = str[i]
	}
	return string(res)
}
