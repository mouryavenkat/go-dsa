package main

func rotatematrix90AntiClock(inputmatrix [][]int, order int, i int, j int) {
	inputmatrix[i][j], inputmatrix[order-1-j][i] = inputmatrix[order-1-j][i], inputmatrix[i][j]
	inputmatrix[i][j], inputmatrix[order-i-1][order-1-j] = inputmatrix[order-i-1][order-1-j], inputmatrix[i][j]
	inputmatrix[i][j], inputmatrix[j][order-i-1] = inputmatrix[j][order-i-1], inputmatrix[i][j]
}
func main() {
	inputmatrix := [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}
	for i := 0; i < len(inputmatrix)/2; i++ {
		for j := i; j < len(inputmatrix)-i-1; j++ {
			rotatematrix90AntiClock(inputmatrix, len(inputmatrix), i, j)
		}
	}
}
