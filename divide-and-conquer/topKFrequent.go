package main
// https://leetcode-cn.com/problems/g5c51o/

func topKFrequent(nums []int, k int) []int {
	size := len(nums)
	if size == 0 {
		return nums
	}
	dict := make(map[int]int)
	for _, num := range nums {
		dict[num]++
	}
	var arr [][]int
	for k, v := range dict {
		arr = append(arr, []int{k, v})
	}
	quickSort(arr, 0, len(arr)-1)
	var res []int
	size = len(arr)
	for i := 1; i <= k; i++ {
		res = append(res, arr[size-i][0])
	}
	return res
}

func quickSort(arr [][]int, left, right int) {
	deal(arr, left, right)
	if right - left < 3 {
		return
	}
	i, j := left+1, right-2
	for {
		for ;arr[i][1] < arr[right-1][1];i++{
		}
		for ;j > i && arr[j][1] >= arr[right-1][1];j--{
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[right-1] = arr[right-1], arr[i]
	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}

func deal(arr [][]int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right - left)/2
	if arr[left][1] > arr[mid][1] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}
	if arr[mid][1] > arr[right][1] {
		arr[mid], arr[right] = arr[right], arr[mid]
	}
	if arr[left][1] > arr[mid][1] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}
	arr[mid], arr[right-1] = arr[right-1], arr[mid]
}
