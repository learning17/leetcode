package main
// https://www.nowcoder.com/practice/9cf027bf54714ad889d4f30ff0ae5481

import (
	"math"
	"fmt"
)
func main() {
	arr := []int{}
	res := LIS(arr)
	fmt.Println(res)
}
func LIS( arr []int ) []int {
	size := len(arr)
	if size < 2 {
		return arr 
	}
	ends := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		ends[i] = math.MaxInt64
	}
	ends[0], ends[1] = 0, arr[0]
	maxSize := 1
	ans := make([][]int, size)
	ans[1] = []int{arr[0]}
	for i := 1; i < size; i++ {
		pos := findBinary(ends, arr[i])
		if ends[pos+1] >= arr[i] {
			ends[pos+1] = arr[i]
			tmp := append([]int{}, ans[pos]...)
			ans[pos+1] = append(tmp, arr[i])
		}
		if pos+1 > maxSize {
			maxSize = pos+1
		}
	}
	return ans[maxSize-1]
}

func findBinary(arr []int, target int) int {
	left, right := 0, len(arr) - 1
	for {
		if left >= right {
			return left
		}
		mid := left + (right-left+1)/2
		if arr[mid] >= target {
			right = mid - 1
		} else {
			left = mid
		}
	}
}
