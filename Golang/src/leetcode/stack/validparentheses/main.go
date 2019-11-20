/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

链接：https://leetcode-cn.com/problems/valid-parentheses
*/
package validparentheses

func isValid(s string) bool {
	var stack []rune
	prthMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, single := range s {
		if match, ok := prthMap[single]; ok {
			if len(stack) > 0 {
				if stack[len(stack)-1] != match {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			} else {
				return false
			}
		} else {
			stack = append(stack, single)
		}
	}
	return len(stack) == 0
}
