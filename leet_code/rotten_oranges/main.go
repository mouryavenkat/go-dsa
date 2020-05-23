package main

func orangesRotting(grid [][]int) int {
	rottenOrangeIndexes := [][]int{}
	freshOrangeCount := 0
	steps := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				rottenOrangeIndexes = append(rottenOrangeIndexes, []int{i, j})
			} else if grid[i][j] == 1 {
				freshOrangeCount += 1
			}
		}
	}

	for len(rottenOrangeIndexes) > 0 {
		newRottenOranges := [][]int{}
		for i := 0; i < len(rottenOrangeIndexes); i++ {
			var x, y int
			rottenIndex := rottenOrangeIndexes[i]
			x = rottenIndex[0]
			y = rottenIndex[1]
			if y+1 < len(grid[x]) && grid[x][y+1] == 1 {
				newRottenOranges = append(newRottenOranges, []int{x, y + 1})
				grid[x][y+1] = 2
				freshOrangeCount--
			}
			if y-1 >= 0 && grid[x][y-1] == 1 {
				newRottenOranges = append(newRottenOranges, []int{x, y - 1})
				grid[x][y-1] = 2
				freshOrangeCount--
			}
			if x+1 < len(grid) && grid[x+1][y] == 1 {
				newRottenOranges = append(newRottenOranges, []int{x + 1, y})
				grid[x+1][y] = 2
				freshOrangeCount--
			}
			if x-1 >= 0 && grid[x-1][y] == 1 {
				newRottenOranges = append(newRottenOranges, []int{x - 1, y})
				grid[x-1][y] = 2
				freshOrangeCount--
			}
		}
		rottenOrangeIndexes = newRottenOranges
		if len(newRottenOranges) > 0 {
			steps++
		}
	}
	if freshOrangeCount > 0 {
		return -1
	}
	return steps
}

func main() {
	/*
		In a given grid, each cell can have one of three values:

		the value 0 representing an empty cell;
		the value 1 representing a fresh orange;
		the value 2 representing a rotten orange.
		Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.

		Return the minimum number of minutes that must elapse until no cell has a fresh orange.  If this is impossible, return -1 instead.
	*/
	orangesRotting([][]int{})
}
