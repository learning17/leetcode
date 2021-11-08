package main
// https://www.nowcoder.com/practice/762836f4d43d43ca9deb273b3de8e1f4

func IsContinuous( numbers []int ) bool {
	var maxNum, minNum, bit int
	for _, num := range numbers {
		if num == 0 {
			continue
		}
		if minNum == 0 {
			minNum = num
		}
		maxNum = max(maxNum, num)
		minNum = min(minNum, num)
		if (bit >> num) & 1 == 1 {
			return false
		} 
		bit |= (1 << num)
	}
	return maxNum - minNum < 5
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

