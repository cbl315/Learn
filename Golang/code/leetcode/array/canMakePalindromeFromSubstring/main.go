/*
Given a string s, we make queries on substrings of s.

For each query queries[i] = [left, right, k], we may rearrange the substring s[left], ..., s[right], and then choose up to k of them to replace with any lowercase English letter.

If the substring is possible to be a palindrome string after the operations above, the result of the query is true. Otherwise, the result is false.

Return an array answer[], where answer[i] is the result of the i-th query queries[i].

Note that: Each letter is counted individually for replacement so if for example s[left..right] = "aaa", and k = 2, we can only replace two of the letters.  (Also, note that the initial string s is never modified by any query.)

链接：https://leetcode-cn.com/problems/can-make-palindrome-from-substring
*/
package canMakePalindromeFromSubstring

func canMakePaliQueries(s string, queries [][]int) []bool {
	/*
		遍历queries 对于每个单独的query_i 计算替换成回文字符串所需的最小替换个数num_i 判断num_i是否不大于给定的个数
	*/
	if len(s) <= 0 || len(queries) <= 0 {
		return []bool{}
	}
	var PosLetterNum = make([][]int, len(s)+1)
	PosLetterNum[0] = make([]int, 26)
	var letterNum = make([]int, 26)
	for i, l := range s {
		letterNum[l-'a']++
		currLetterNum := make([]int, 26)
		copy(currLetterNum, letterNum)
		PosLetterNum[i+1] = currLetterNum
	}
	var res = make([]bool, len(queries))
	for i, query := range queries {
		res[i] = false
		left, right, k := query[0], query[1], query[2]
		maxK := (right - left + 1) / 2
		if k >= maxK || k >= 13 {
			res[i] = true
			continue
		}
		// rearrange
		leftLetterNum, rightLetterNum := PosLetterNum[left], PosLetterNum[right+1]
		substrLetterNum := make([]int, 26)
		for j, num := range leftLetterNum {
			substrLetterNum[j] = rightLetterNum[j] - num
		}
		for _, num := range substrLetterNum {
			maxK -= num / 2
		}
		if k >= maxK {
			res[i] = true
		}
	}
	return res
}
