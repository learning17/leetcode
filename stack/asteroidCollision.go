package main
// https://leetcode-cn.com/problems/XagZNi/
// 小行星碰撞
func asteroidCollision(asteroids []int) []int {
	var stack1, stack2 []int
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			stack1 = append(stack1, asteroids[i])
			continue
		}
		for len(stack1) > 0 && stack1[len(stack1)-1] < -asteroids[i] {
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack1) == 0 {
			stack2 = append(stack2, asteroids[i])
		} else if stack1[len(stack1)-1] == -asteroids[i] {
			stack1 = stack1[:len(stack1)-1]
		}
	}
	return append(stack2, stack1...)
}

