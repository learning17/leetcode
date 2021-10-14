package main
// https://www.nowcoder.com/practice/e016ad9b7f0b45048c58a9f27ba618bf

func findKth( a []int ,  n int ,  K int ) int {
	quickSort(a, 0, n-1)
	return a[n-K]
}

func quickSort(arr []int, left, right int) {
	help(arr, left, right)
	if right - left < 3 {
		return
	}
	i, j := left+1, right-2
	for {
		for ;arr[i] < arr[right-1];i++ {
		}
		for ;j > i && arr[j] >= arr[right-1];j--{
		}
		if j <= i {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[right-1] = arr[right-1], arr[i]
	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}

func help(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right- left)/2
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
