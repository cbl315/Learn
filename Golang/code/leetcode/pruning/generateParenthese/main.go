/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

链接：https://leetcode-cn.com/problems/generate-parentheses
*/
package generateParenthese

func generateParenthesis(n int) []string {
	return gen(0, 0, n, "", []string{})
}

func gen(leftUsed, rightUsed, n int, currResult string, currList []string) []string {
	if leftUsed == n && rightUsed == n {
		currList = append(currList, currResult)
		return currList
	}
	if leftUsed < n {
		currList = gen(leftUsed+1, rightUsed, n, currResult+"(", currList)
	}
	if rightUsed < n && rightUsed < leftUsed {
		currList = gen(leftUsed, rightUsed+1, n, currResult+")", currList)
	}
	return currList
}
