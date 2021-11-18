package main
// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
// 替换空格


func replaceSpace(s string) string {
	var bytes []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			bytes = append(bytes, []byte{'%', '2', '0'}...)
		} else {
			bytes = append(bytes, s[i])
		}
	}
	return string(bytes)
}

