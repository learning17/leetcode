package main
// https://www.nowcoder.com/practice/28eb3175488f4434a4a6207f6f484f47

func longestCommonPrefix( strs []string ) string {
	size := len(strs)
	if size == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < size; i++ {
		prefixSize := len(prefix)
		if prefixSize == 0 {
			return prefix
		}
		tmp := []byte{}
		for j := 0; j < prefixSize && j < len(strs[i]); j++ {
			if prefix[j] != strs[i][j] {
				break
			}
			tmp = append(tmp, prefix[j])
		}
		prefix = string(tmp)
	}
	return prefix
}
