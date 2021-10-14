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
	for k := 2*i+1; k < size; k = 2*k+1 {
		if k+1 < size && arr[k] < arr[k+1] {
			k++
		}
		if arr[i] < arr[k] {
			arr[i], arr[k] = arr[k], arr[i]
			i = k
		} else {
			break
		}
	}
}
