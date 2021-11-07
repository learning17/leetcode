package main
// https://www.nowcoder.com/practice/ef1f53ef31ca408cada5093c8780f44b

func reOrderArray( array []int ) []int {
	var even, odd []int
	for _, a := range array {
		if a % 2 == 0 {
			even = append(even, a)
		} else {
			odd = append(odd, a)
		}
	}
	return append(odd, even...)
}

