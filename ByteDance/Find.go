package main
// https://www.nowcoder.com/practice/abc3fe2ce8e146608e868a70efebf62e

func Find( target int ,  array [][]int ) bool {
	row_size := len(array)
	if row_size == 0 {
		return false
	}
	col_size := len(array[0])

	var check func(r1, r2, c1, c2 int) bool
	check = func(r1, r2, c1, c2 int) bool {
		if r1 < 0 || r1 > r2 || r2 >= row_size || c1 < 0 || c2 >= col_size || c1 > c2 {
			return false
		}
		rm, cm := r1+(r2-r1)/2, c1+(c2-c1)/2
		if array[rm][cm] == target {
			return true
		} else if array[rm][cm] > target {
			return check(r1, r2, c1, cm-1) || check(r1, rm-1, cm, c2)
		} else {
			return check(r1, r2, cm+1, c2) || check(rm+1, r2, c1, cm)
		}
	}
	return check(0, row_size-1, 0, col_size-1)
}

