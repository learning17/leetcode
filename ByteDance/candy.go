package main
// https://www.nowcoder.com/practice/76039109dd0b47e994c08d8319faa352

func candy( arr []int ) int {
	size := len(arr)
	if size == 0 {
		return 0
	}
	candys := make([]int, size)
	for i := 0; i < size; i++ {
		candys[i] = 1
		if i > 0 && arr[i] > arr[i-1] {
			candys[i] = candys[i-1]+1
		}
	}
	for i := size - 1; i > 0; i-- {
		if arr[i-1] > arr[i] && candys[i-1] <= candys[i] {
			candys[i-1] = candys[i]+1
		}
	}
	sum := 0
	for i := 0; i < size; i++ {
		sum += candys[i]
	}
	return sum
}

