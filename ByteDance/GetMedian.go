package main

var arr []int
func Insert(num int){
	if len(arr) == 0 {
		arr = append(arr, num)
		return
	}
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] <= num {
			left = mid+1
		} else {
			right = mid
		}
	}
	if arr[left] > num {
		arr = append(arr[:left], append([]int{num}, arr[left:]...)...)
	} else {
		arr = append(arr[:left+1], append([]int{num}, arr[left+1:]...)...)
	}
}

func GetMedian() float64{
	size := len(arr)
	if size % 2 == 0 {
		return (float64(arr[size/2]) + float64(arr[size/2-1]))/2
	}
	return float64(arr[size/2])
}

