package main
// https://www.nowcoder.com/practice/ae25fb47d34144a08a0f8ff67e8e7fb5

func foundMonotoneStack( nums []int ) [][]int {
	size := len(nums)
	if size == 0 {
		return [][]int{}
	}
	var stack [][]int
	pos := make([][]int, size)
	for i := 0; i < size; i++ {
		pos[i] = make([]int, 2)
	}
	pos[0][0] = -1
	pos[size-1][1] = -1
	stack = append(stack, []int{nums[0], 0})
	for i := 1; i < size; i++ {
		for len(stack) > 0 && nums[i] <= stack[len(stack)-1][0] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			pos[i][0] = -1
		} else {
			pos[i][0] = stack[len(stack)-1][1]
		}
		stack = append(stack, []int{nums[i], i})
	}
	stack = stack[:0]
	stack = append(stack, []int{nums[size-1], size-1})
	for i := size - 2; i >= 0; i-- {
		for len(stack) > 0 && nums[i] <= stack[len(stack)-1][0] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			pos[i][1] = -1
		} else {
			pos[i][1] = stack[len(stack)-1][1]
		}
		stack = append(stack, []int{nums[i], i})
	}
	return pos
}

