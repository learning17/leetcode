package main
// https://www.nowcoder.com/practice/55fb3c68d08d46119f76ae2df7566880
import (
	"strings"
	"strconv"
)

func solve( IP string ) string {
	if isIPV4(IP) {
		return "IPv4"
	}
	if isIPV6(IP) {
		return "IPv6"
	}
	return "Neither"
}

func isIPV4(IP string) bool {
	arr := strings.Split(IP, ".")
	if len(arr) != 4 {
		return false
	}
	for i := 0; i < 4; i++ {
		v, err := strconv.Atoi(arr[i])
		if err != nil {
			return false
		}
		if v < 0 || v > 255 {
			return false
		}
		if v == 0 && len(arr[i]) > 1 {
			return false
		}
		if v > 0 && arr[i][0] == '0' {
			return false
		}
	}
	return true
}

func isIPV6(IP string) bool {
	arr := strings.Split(IP, ":")
	if len(arr) != 8 {
		return false
	} 
	for i := 0; i < 8; i++ {
		if len(arr[i]) > 4 {
			return false
		}
		for j := 0; j < len(arr[i]); j++ {
			if (arr[i][j] >= '0' && arr[i][j] <= '9') || 
				(arr[i][j] >= 'a' && arr[i][j] <= 'f') || 
				(arr[i][j] >= 'A' && arr[i][j] <= 'F') {
					continue
				}
				return false
		}
	}
	return true
}

