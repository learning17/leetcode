package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	d := make(map[byte]bool)
	i, j, maxSize := 0, 0, 0
	for ; i <= j && j < len(s); {
		flag, ok := d[s[j]]
		if !flag || !ok {
			d[s[j]] = true
			maxSize = max(maxSize, j-i+1)
			j++
		} else {
			d[s[i]] = false
			i++
		}
	}
	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

