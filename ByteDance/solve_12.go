package main
// https://www.nowcoder.com/practice/e19927a8fd5d477794dac67096862042

func solve( n int ,  m int ,  a []int ) []int {
	if m > n {
		m %= n
	}
	reverse(a)
	reverse(a[:m])
	reverse(a[m:])
	return a
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

