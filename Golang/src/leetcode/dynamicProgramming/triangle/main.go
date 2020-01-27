/*
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:

Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

链接：https://leetcode-cn.com/problems/triangle
*/
package triangle

func MyMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minimumTotal(triangle [][]int) int {
	// DP[i, j] = min(DP[i+1, j], DP[i+1, j+1]) + triangle[i, j]
	if triangle == nil || len(triangle) == 0 {
		return 0
	}
	var distance = make(map[int]map[int]int, len(triangle)+1)
	distance[len(triangle)] = make(map[int]int, len(triangle)+1)
	for i := len(triangle) - 1; i >= 0; i-- {
		distance[i] = make(map[int]int, i+1)
		for j := i; j >= 0; j-- {
			distance[i][j] = triangle[i][j] + MyMin(distance[i+1][j], distance[i+1][j+1])
		}
	}
	return distance[0][0]
}
