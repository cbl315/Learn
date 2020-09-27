/*
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

链接：https://leetcode-cn.com/problems/number-of-islands
*/
package numberOfIslands

import "fmt"

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}
	inst := FloodFill{
		Grid:    grid,
		MaxRow:  len(grid),
		MaxCol:  len(grid[0]),
		Visited: map[string]bool{},
	}
	return inst.FloodFillSolution()
}

/*
染色法：
遍历矩阵 如果非陆地（即为0） 跳过; 如果是陆地，count++ 并将周围的所有陆地（包括不直接相邻的陆地）置为1
*/

type FloodFill struct {
	Grid    [][]byte
	MaxRow  int
	MaxCol  int
	Visited map[string]bool
}

func (f *FloodFill) FloodFillSolution() int {
	count := 0
	for rowIndex, row := range f.Grid {
		for colIndex, _ := range row {
			count += f.floodFillDFS(rowIndex, colIndex)
		}
	}
	return count
}

func (f *FloodFill) floodFillDFS(x, y int) int {
	if !f.isValid(x, y) {
		return 0
	}
	index := fmt.Sprintf("%d-%d", x, y)
	f.Visited[index] = true
	f.floodFillDFS(x+1, y)
	f.floodFillDFS(x-1, y)
	f.floodFillDFS(x, y+1)
	f.floodFillDFS(x, y-1)
	return 1
}

func (f *FloodFill) isValid(x, y int) bool {
	if x < 0 || x >= f.MaxRow || y < 0 || y >= f.MaxCol {
		return false
	}
	index := fmt.Sprintf("%d-%d", x, y)
	if _, ok := f.Visited[index]; ok {
		return false
	}
	if f.Grid[x][y] == '0' {
		return false
	}
	return true
}

/*
 并查集法
*/
func numIslands2(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}
	return uf(grid)
}
func uf(grid [][]byte) int {
	// init uf
	uf := unionFind{}
	var ufIndex, maxCol, maxRow = []int{}, len(grid[0]), len(grid)
	for rowIndex, row := range grid {
		for colIndex, land := range row {
			index := rowIndex*maxCol + colIndex
			if land == '1' {
				ufIndex = append(ufIndex, index)
			}
		}
	}
	uf.Init(ufIndex)
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	// union
	for rowIndex, row := range grid {
		for colIndex, land := range row {
			index := rowIndex*maxCol + colIndex
			// find whether left or top land exist
			if land == '1' {
				grid[rowIndex][colIndex] = '0'
				for i := 0; i < len(dx); i++ {
					nearRow, nearCol := rowIndex+dx[i], colIndex+dy[i]
					if nearRow >= 0 && nearRow < maxRow && nearCol >= 0 && nearCol < maxCol && grid[nearRow][nearCol] == '1' {
						nearIndex := nearRow*maxCol + nearCol
						uf.Union(index, nearIndex)
					}
				}
			}
		}
	}
	return uf.Count()
}
