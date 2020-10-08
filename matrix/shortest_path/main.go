package shortest_path

/*
	In an N by N square grid, each cell is either empty (0) or blocked (1).

	A clear path from top-left to bottom-right has length k if and only if it is composed of cells C_1, C_2, ..., C_k such that:

	Adjacent cells C_i and C_{i+1} are connected 8-directionally (ie., they are different and share an edge or corner)
	C_1 is at location (0, 0) (ie. has value grid[0][0])
	C_k is at location (N-1, N-1) (ie. has value grid[N-1][N-1])
	If C_i is located at (r, c), then grid[r][c] is empty (ie. grid[r][c] == 0).
	Return the length of the shortest such clear path from top-left to bottom-right.  If such a path does not exist, return -1.


*/
func updateStackAndReturnSmallestPath(grid [][]int, stack [][2]int) ([][2]int, int) {
	element := stack[0]
	gridLen := len(grid) * len(grid)
	var smallest = gridLen
	for i := element[0] - 1; i <= element[0]+1; i++ {
		for j := element[1] - 1; j <= element[1]+1; j++ {
			if i < 0 || i >= len(grid) || j < 0 || j >= len(grid) || grid[i][j] == -1 {
				continue
			}
			if grid[i][j] == 0 {
				stack = append(stack, [2]int{i, j})
				grid[i][j] = -1
				continue
			}
			if grid[i][j] > 0 && grid[i][j] < smallest {
				smallest = grid[i][j]
			}
		}
	}
	if smallest == gridLen {
		return stack, 0
	}
	return stack, smallest
}
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[len(grid)-1][len(grid)-1] == 1 || grid[0][0] == 1 {
		return -1
	}
	gridLen := len(grid)
	var stack [][2]int
	for i := 0; i < gridLen; i++ {
		for j := 0; j < gridLen; j++ {
			grid[i][j] = -grid[i][j]
		}
	}
	grid[0][0] = -1
	stack = [][2]int{{0, 0}}
	for len(stack) > 0 {
		var smallest int
		stack, smallest = updateStackAndReturnSmallestPath(grid, stack)
		poppedNode := stack[0]
		stack = stack[1:]
		grid[poppedNode[0]][poppedNode[1]] = smallest + 1
	}
	if grid[gridLen-1][gridLen-1] == 0 {
		return -1
	}
	return grid[gridLen-1][gridLen-1]
}
