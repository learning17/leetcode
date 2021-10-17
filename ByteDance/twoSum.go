package main
// https://www.nowcoder.com/practice/20ef0972485e41019e39543e8e895b7f

import "sort"

func main() {
	numbers := []int{3,2,4}
	twoSum(numbers, 6)
}

func twoSum( numbers []int ,  target int ) []int {
	size := len(numbers)
	if size == 0 {
		return []int{}
	}
	arr := make([][]int, size)
	for i := 0; i < size; i++ {
		arr[i] = []int{i, numbers[i]}
	}
	sort.Slice(arr, func(i, j int) bool {return arr[i][1] < arr[j][1]})
	for i, j := 0, size-1; i < size && j >= 0; {
		sum := arr[i][1] + arr[j][1]
		if sum == target {
			return posSort(arr[i][0]+1, arr[j][0]+1)
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{-1,-1}
}

func posSort(a, b int) []int {
	if a < b {
		return []int{a, b}
	}
	return []int{b, a}
}
