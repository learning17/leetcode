package main
// https://www.nowcoder.com/practice/5d3d74c3bf7f4e368e03096bb8857871

func foundOnceNumber( arr []int ,  k int ) int {
	var sum int
	for i := 31; i >= 0; i-- {
		var cnt int
		for j := 0; j < len(arr); j++ {
			cnt += (arr[j]>>i)&1
		}
		sum = 2*sum + cnt%k
	}
	if (sum >> 31) & 1 == 0 {
		return sum
	}
	// 负数反码表示
	sum &= ((1 << 31)-1)
	sum--
	for i := 0; i < 31; i++ {
		bit := 1 << i
		sum ^= bit
	}
	return -sum
}
