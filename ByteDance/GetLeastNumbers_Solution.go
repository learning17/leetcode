package main

func GetLeastNumbers_Solution( input []int ,  k int ) []int {
	for i := (k-2)/2; i >= 0; i-- {
		adjust(input, i, k)
	}
	for i := k; i < len(input); i++ {
		if input[i] >= input[0] {
			continue
		}
		input[0] = input[i]
		adjust(input, 0, k)
	}
	return input[:k]
}

func adjust(arr []int, i, size int) {
	for {
		left, right := 2*i+1, 2*i+2
		if left >= size {
			break
		}
		if right >= size {
			if arr[left] < arr[i] {
				break
			}
			arr[left], arr[i] = arr[i], arr[left]
			i = left
		} else {
			if arr[left] <= arr[i] && arr[right] <= arr[i] {
				break
			}
			if arr[left] > arr[right] {
				arr[left], arr[i] = arr[i], arr[left]
				i = left
			} else {
				arr[right], arr[i] = arr[i], arr[right]
				i = right
			}
		}
	}
}
