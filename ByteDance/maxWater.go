package main
// https://www.nowcoder.com/practice/31c1aed01b394f0b8b7734de0324e00f

func maxWater( arr []int ) int64 {
	size := len(arr)
	if size < 3 {
		return 0
	}
	leftMax, rightMax := arr[0], arr[size-1]
	sum := 0
	for left, right := 1, size-2; left <= right; {
		leftMax, rightMax = max(leftMax, arr[left]), max(rightMax, arr[right])
		if leftMax < rightMax {
			sum += leftMax - arr[left]
			left++
		} else {
			sum += rightMax - arr[right]
			right--
		}
	}
	return int64(sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

