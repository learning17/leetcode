package main
// https://www.nowcoder.com/practice/389fc1c3d3be4479a154f63f495abff8

func FindNumsAppearOnce( array []int ) []int {
	var bit, a, b int
	for _, num := range array {
		bit ^= num
	}
	bit &=-bit
	for _, num := range array {
		if num & bit == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	if a > b {
		return []int{b, a}
	}
	return []int{a, b}
}
