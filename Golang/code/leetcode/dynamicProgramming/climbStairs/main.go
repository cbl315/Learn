/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

链接：https://leetcode-cn.com/problems/climbing-stairs
*/
package climbStairs

func climbStairs(n int) int {
	var fibo = make(map[int]int, n+1)
	fibo[0], fibo[1] = 1, 1
	for i := 2; i <= n; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
	}
	return fibo[n]
}
