package main
import "sort"

func permutation(S string) []string {
    arr := []byte(S)
	sort.Slice(arr, func(i, j int) bool {return arr[i] < arr[j]})
    var path []byte
    var res []string
    size := len(S)
    if size == 0 {
        return res
    }
    var backtrack func([]byte)
    backtrack = func(subArr []byte) {
        if len(path) == size {
            res = append(res, string(path))
            return
        }
        for i, a := range subArr {
			if i > 0 && subArr[i] == subArr[i-1] {
				continue
			}
            path = append(path, a)
            tmp := append([]byte{}, subArr[:i]...)
            tmp = append(tmp, subArr[i+1:]...)
            backtrack(tmp)
            path = path[:len(path)-1]
        }
    }
    backtrack(arr)
    return res
}
