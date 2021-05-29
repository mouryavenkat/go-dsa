package main

/*
	The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

	Given an integer n, return the number of distinct solutions to the n-queens puzzle.

*/

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getPossibleSlots(n int, column int, occupiedSlots [][]int) [][]int {
	var validSlots [][]int
	for i := 0; i < n; i++ {
		isValidSlot := true
		for _, occupiedSlot := range occupiedSlots {
			if occupiedSlot[0] == i {
				isValidSlot = false
				break
			}
			if abs(occupiedSlot[0]-i) == abs(occupiedSlot[1]-column) {
				isValidSlot = false
				break
			}
		}
		if isValidSlot {
			validSlots = append(validSlots, []int{i, column})
		}
	}
	return validSlots
}

func totalNQueens(n int) int {
	var count int
	var occupiedSlots [][]int
	possibleSlotsPerIndex := make([][][]int, n)
	var i int
	for i < n {
		var possibleSlots [][]int
		if len(possibleSlotsPerIndex[i]) == 0 {
			possibleSlots = getPossibleSlots(n, i, occupiedSlots)
		} else {
			possibleSlots = possibleSlotsPerIndex[i]
		}
		if len(possibleSlots) == 0 || i == n-1 {
			if i == n-1 && len(possibleSlots) > 0 {
				count++
			}
			for i >= 0 && len(possibleSlotsPerIndex[i]) == 0 {
				if i > 0 {
					occupiedSlots = occupiedSlots[0 : i-1]
				}
				i--
			}
			if i < 0 {
				return count
			}
			continue
		}
		occupiedSlots = append(occupiedSlots, possibleSlots[0])
		if len(possibleSlots) > 1 {
			possibleSlotsPerIndex[i] = possibleSlots[1:]
		} else {
			possibleSlotsPerIndex[i] = [][]int{}
		}
		i++
	}
	return 0
}
