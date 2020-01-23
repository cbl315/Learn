/*
Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the Hamming weight).

链接：https://leetcode-cn.com/problems/number-of-1-bits
*/
package numOf1bit

func hammingWeight(num uint32) int {
	var cnt int
	var mask uint32 = 1
	//for num >= mask{  // 这里不能这样判断 mask可能会超过32位的限制 进位到33位 第33位被舍弃 从而导致mask等于0
	for i := 0; i < 32; i++ {
		if num&mask > 0 {
			cnt++
		}
		mask = mask << 1
	}
	return cnt
}

func hammingWeight2(num uint32) int {
	var cnt = 0
	for num != 0 {
		cnt++
		num = num & (num - 1)
	}
	return cnt
}
