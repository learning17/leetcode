package main

// https://leetcode-cn.com/problems/shortest-supersequence-lcci/

func shortestSeq(big []int, small []int) []int {
	big_size, small_size := len(big), len(small)
	if big_size < small_size || small_size == 0 {
		return []int{}
	}
	dict := make(map[int]int)
	for i := 0; i < small_size; i++ {
		dict[big[i]]++
	}
	minSize, minLeft, minRight := big_size, -1, -1
	for left, right := 0, small_size-1; left <= right && right < big_size; {
		if isValid(small, dict) {
			if right - left < minSize {
				minLeft, minRight, minSize = left, right, right - left
			}
			left++
			dict[big[left]]--
		} else {
			right++
			if right >= big_size {
                break
            }
			dict[big[right]]++
		}
	}
	if minLeft != -1 {
		return []int{minLeft, minRight}
	} else {
		return []int{}
	}
}

func isValid(small []int, dict map[int]int) bool {
	for _, s := range small {
		if tmp, ok := dict[s]; !ok || tmp <= 0 {
			return false
		}
	}
	return true
}

