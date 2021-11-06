package main
// https://www.nowcoder.com/practice/96bd6684e04a44eb80e6a68efc0ec6c5

func InversePairs( data []int ) int {
	if len(data) < 2 {
		return 0
	}

	var merge func(int, int) int
	merge = func(left, right int) int {
		if left >= right {
			return 0
		}
		mid := left + (right-left)/2
		num := merge(left, mid) + merge(mid+1, right)
		for i := mid+1; i <= right; i++ {
			for j := left; j <= mid; j++ {
				if data[j] > data[i] {
					num++
				}
			}
		}
		return num
	}
	return merge(0, len(data)-1) % 1000000007
}

