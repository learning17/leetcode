package main
// https://www.nowcoder.com/practice/6fbe70f3a51d44fa9395cfc49694404f

func findMedianinTwoSortedAray( arr1 []int ,  arr2 []int ) int {
	size := len(arr1)
	var ans int
	for i, j, k := 0, 0, 0; i < size && j < size; k++ {
		if k == size - 1 {
			ans = min(arr1[i], arr2[j])
			break
		}
		if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
