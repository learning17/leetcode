package main
// https://www.nowcoder.com/practice/e8a1b01a2df14cb2b228b30ee6a92163

func MoreThanHalfNum_Solution( numbers []int ) int {
	var res, num int
	for i := 0; i < len(numbers); i++ {
		if num == 0 || res == numbers[i] {
			res = numbers[i]
			num++
		} else {
			num--
		}
	}
	return res
}

