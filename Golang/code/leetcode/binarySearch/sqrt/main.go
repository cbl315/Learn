/*
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

链接：https://leetcode-cn.com/problems/sqrtx
*/
package sqrt

import "math"

// 二分法
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 1, x
	var res int
	for left <= right {
		mid := (left + right) / 2
		currX := x / mid
		if currX == mid {
			return mid
		} else if mid > currX {
			right = mid - 1
		} else {
			left = mid + 1
			res = mid
		}
	}
	return res
}

// 二分法
// x >= 0
func mySqrt2(x float64) float64 {
	left, right := 0.0, x
	fabs := 1e-8
	for left <= right {
		mid := (left + right) / 2.0
		currX := mid * mid
		if math.Abs(currX-x) <= fabs {
			return mid
		} else if x > currX {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
