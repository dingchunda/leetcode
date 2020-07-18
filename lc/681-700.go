package lc

import "container/list"

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	tiles := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		tiles[i] = make([]bool, len(grid[0]))
	}
	type tile struct {
		i, j int
		land int
	}
	var high = list.New()
	high.PushBack(tile{0, 0, -1})

	area := 0
	maxArea := 0
	for high.Len() > 0 {
		var next tile
		next = high.Front().Value.(tile)
		high.Remove(high.Front())
		i, j := next.i, next.j
		land := next.land
		if tiles[i][j] {
			continue
		}
		tiles[i][j] = true
		ty := grid[i][j]
		if ((ty == 0 || ty == 1) && land != 1) && area > 0 {
			if area > maxArea {
				maxArea = area
			}
			area = 0
		}
		if ty == 1 && land != 1 {
			area = 1
		}
		if ty == 1 && land == 1 {
			area++
		}
		if i-1 >= 0 && !tiles[i-1][j] {
			if grid[i-1][j] == 1 {
				high.PushFront(tile{i - 1, j, grid[i][j]})
			} else {
				high.PushBack(tile{i - 1, j, grid[i][j]})
			}
		}
		if i+1 < len(grid) && !tiles[i+1][j] {
			if grid[i+1][j] == 1 {
				high.PushFront(tile{i + 1, j, grid[i][j]})
			} else {
				high.PushBack(tile{i + 1, j, grid[i][j]})
			}
		}
		if j-1 >= 0 && !tiles[i][j-1] {
			if grid[i][j-1] == 1 {
				high.PushFront(tile{i, j - 1, grid[i][j]})
			} else {
				high.PushBack(tile{i, j - 1, grid[i][j]})
			}
		}
		if j+1 < len(grid[0]) && !tiles[i][j+1] {
			if grid[i][j+1] == 1 {
				high.PushFront(tile{i, j + 1, grid[i][j]})
			} else {
				high.PushBack(tile{i, j + 1, grid[i][j]})
			}
		}
	}
	if area > 0 && area > maxArea {
		maxArea = area
	}
	return maxArea
}
