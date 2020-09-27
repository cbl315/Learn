/*
Given two strings s and t , write a function to determine if t is an anagram of s.

链接：https://leetcode-cn.com/problems/valid-anagram
*/
package validAnagram

func isAnagram(s string, t string) bool {
	var cnt = make(map[rune]int)
	var cnt2 = make(map[rune]int)
	for _, let := range s {
		cnt[let]++
	}
	for _, let := range t {
		cnt2[let]++
	}
	if len(cnt) != len(cnt2) {
		return false
	}
	for k, v := range cnt {
		if cnt2[k] != v {
			return false
		}
	}
	return true
}
