package main
// https://leetcode-cn.com/problems/aseY1I/

func maxProduct(words []string) int {
	size := len(words)
	if size < 2 {
		return 0
	}
	var wordsInt []int
	dict := make(map[string]int)
	for _, word := range words {
		wordsInt = append(wordsInt, wordToInt(word))
		dict[word] = len(word)
	}
	maxValue := 0
	for i := 0; i < size; i++ {
		for j := i+1; j < size; j++ {
			if wordsInt[i] & wordsInt[j] != 0 {
				continue
			}
			tmp := dict[words[i]] * dict[words[j]]
			if tmp > maxValue {
				maxValue = tmp
			}
		}
	}
	return maxValue
}

func wordToInt(word string) int {
	ans := 0
	for _,c := range word {
		pos := c - 'a'
		ans |= (1 << pos)
	}
	return ans
}
