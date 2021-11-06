package main
// https://www.nowcoder.com/practice/c466d480d20c4c7c9d322d12ca7955ac

func minWindow( S string ,  T string ) string {
	sizeS, sizeT := len(S), len(T)
	if sizeS == 0 || sizeS < sizeT {
		return ""
	}
	dictT := make(map[byte]int)
	for i := 0; i < sizeT; i++ {
		dictT[T[i]]++
	}
	dict := make(map[byte]int)
	check := func() bool {
		for k,v := range dictT {
			if tmpV, ok := dict[k]; !ok || tmpV < v {
				return false
			}
		}
		return true
	}
	minSize, ans := sizeS, ""
	for left, right := 0, 0; left <= right && right < sizeS; {
		dict[S[right]]++
		right++
		for check() {
			if minSize >= right - left {
				minSize = right - left
				ans = S[left:right]
			}
			dict[S[left]]--
			left++
		}
	}
	return ans
}

