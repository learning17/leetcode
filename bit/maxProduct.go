package main
// https://leetcode-cn.com/problems/aseY1I/
// 单词长度的最大乘积

func maxProduct(words []string) int {
	size := len(words)
	if size < 2 {
		return 0
	}
	dict := make(map[string][]int)
	for _, word := range words {
		dict[word] = []int{wordToInt(word), len(word)}
	}
	var maxSize int
	for i := 0; i < size; i++ {
		for j := i+1; j < size; j++ {
			if dict[words[i]][0] & dict[words[j]][0] != 0{
				continue
			}
			maxSize = max(maxSize, dict[words[i]][1] * dict[words[j]][1]) 
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

func wordToInt(word string) int {
	var ans int
	for _,c := range word {
		ans |= (1 << int(c-'a'))
	}
	return ans
}

