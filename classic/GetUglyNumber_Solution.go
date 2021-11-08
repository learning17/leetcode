package main
// https://www.nowcoder.com/practice/6aa9e04fc3794f68acf8778237ba065b

func GetUglyNumber_Solution( index int ) int {
	if index == 0 {
		return 0
	}
	arr := make([]int, index)
	arr[0] = 1
	var x, y, z int
	for i := 1; i < index; i++ {
		arr[i] = min(2*arr[x], 3*arr[y], 5*arr[z])
		if arr[i] == 2*arr[x] {
			x++
		}
		if arr[i] == 3*arr[y] {
			y++
		}
		if arr[i] == 5*arr[z] {
			z++
		}
	}
	return arr[index-1]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

