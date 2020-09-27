/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数。
*/
package pow

import (
	"math"
)

func myPow(x float64, n int) float64 {
	// divide and conquer
	if n == 0 {
		return 1
	} // x^0=1,
	if n < 0 {
		return 1 / myPow(x, -n)
	} // x^n=1/x^-n , n<0
	if n%2 == 0 {
		return myPow(x*x, n/2)
	} else {
		return x * myPow(x, (n-1))
	}
}

func myPow2(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

func myPow3(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow3(x, -n)
	}
	result := 1.0
	for n > 0 {
		if uint(n)&1 == 1 {
			result *= x
		}
		x *= x
		n = n >> 1
	}
	return result
}
