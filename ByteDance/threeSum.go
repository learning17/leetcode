package main
// https://www.nowcoder.com/practice/345e2ed5f81d4017bbb8cc6055b0b711

import (
	"sort"
)

func threeSum( num []int ) [][]int {
	sort.Ints(num)
	ans := [][]int{}
	size := len(num)
	if size < 3 {
		return ans
	}
	for i := 0; i < size - 2; i++ {
		if num[i] > 0 {
			break
		}
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		left, right := i+1, size-1
		for left < right {
			sum := num[i] + num[left] + num[right]
			if sum == 0 {
				ans = append(ans, []int{num[i], num[left], num[right]})
				left++
				right--
				for left < right && num[left-1] == num[left]{
					left++
				}
				for left < right && num[right] == num[right+1]{
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

