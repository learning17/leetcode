package main
// https://www.nowcoder.com/practice/cf4091ca75ca47958182dae85369c82c

func gcd( a int ,  b int ) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

