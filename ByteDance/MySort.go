package main
// https://www.nowcoder.com/practice/2baf799ea0594abd974d37139de27896

func MySort( arr []int ) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func heapSort(arr []int) {
	size := len(arr)
	if size <= 1 {
		return
	}
	for k := (size-2)/2; k >= 0; k-- {
		adjust(arr, k, size)
	}
	for i := size - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		adjust(arr, 0, i)
	}
}

func adjust(arr []int, k, size int) {
	for i := 2*k+1; i < size; i, k = 2*i+1, i {
		if i + 1 < size && arr[i] < arr[i+1] {
			i++
		}
		if arr[i] <= arr[k] {
			break
		}
		arr[i], arr[k] = arr[k], arr[i]
	}
}

// 快排
func quickSort(arr [] int, left, right int) {
	deal(arr, left, right)
	if right - left < 3 {
		return
	}
	i, j := left+1, right-2
	for {
		for ;arr[i] < arr[right-1];i++{
		}
		for ;i < j && arr[j] >= arr[right-1];j--{
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

func deal(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right - left)/2
	if arr[left] > arr[mid] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}
	if arr[mid] > arr[right] {
		arr[mid], arr[right] = arr[right], arr[mid]
	}
	if arr[left] > arr[mid] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}
	arr[mid], arr[right-1] = arr[right-1], arr[mid]
}

