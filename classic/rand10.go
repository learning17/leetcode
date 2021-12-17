package main
// https://leetcode-cn.com/problems/implement-rand10-using-rand7/
/*
用 Rand7() 实现 Rand10()
(randX-1)*Y+randY 可以等概率生成[1, XY] 的随机数
*/
import (
	"math/rand"
	"time"
)

func rand10() int {
	for {
		idx := (rand7()-1)*7 + rand7()
		if idx <= 40 {
			return 1 + (idx-1) % 10
		}
	}
}

func rand7() int {
	rand.Seed(int64(time.Millisecond))
	return rand.Intn(7)
}

