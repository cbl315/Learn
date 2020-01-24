/*
Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.

Example 1:

Input: 2
Output: [0,1,1]

链接：https://leetcode-cn.com/problems/counting-bits
*/
package countingBits

func countBits(num int) []int {
	var bitNums = make(map[int]int, num+1)
	var result = make([]int, num+1)
	for i := 1; i <= num; i++ {
		bitNums[i] = bitNums[i&(i-1)] + 1
		result[i] = bitNums[i]
	}
	return result
}
