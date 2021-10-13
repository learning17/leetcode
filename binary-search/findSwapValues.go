package main

//https://leetcode-cn.com/problems/sum-swap-lcci/

import (
	"sort"
)

func findSwapValues(array1 []int, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	gap := sum(array1) - sum(array2)
	if gap % 2 != 0 {
		return []int{}
	}
	for i, j := 0, 0; i < len(array1) && j < len(array2); {
		tmpGap := 2*(array1[i] - array2[j])
		if tmpGap == gap {
			return []int{array1[i], array2[j]}
		}
		if tmpGap > gap {
			j++
		} else {
			i++
		}
	} 
	return []int{}
}

func sum(arr []int) int {
	var tmpSum int
	for i := 0; i < len(arr); i++ {
		tmpSum += arr[i]
	}
	return tmpSum
}
