/*
Given an integer, write a function to determine if it is a power of two.

Example 1:

Input: 1
Output: true
Explanation: 20 = 1

链接：https://leetcode-cn.com/problems/power-of-two
*/
package powerOfTwo

func isPowerOfTwo(n int) bool {
	return n > 0 && (n-1)&n == 0
}
