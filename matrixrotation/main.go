package main

import "fmt"

func rotatematrix(inputmatrix [][]int, order int, inner int) {
	for i := inner; i < order-inner-1; i++ {
		inputmatrix[inner][inner], inputmatrix[inner][i+1] = inputmatrix[inner][i+1], inputmatrix[inner][inner]
	}
	for i := inner + 1; i < order-inner; i++ {
		inputmatrix[inner][inner], inputmatrix[i][order-inner-1] = inputmatrix[i][order-inner-1], inputmatrix[inner][inner]
	}
	for i := order - inner - 2; i >= inner; i-- {
		inputmatrix[inner][inner], inputmatrix[order-inner-1][i] = inputmatrix[order-inner-1][i], inputmatrix[inner][inner]
	}
	for i := order - inner - 2; i > inner; i-- {
		inputmatrix[inner][inner], inputmatrix[i][inner] = inputmatrix[i][inner], inputmatrix[inner][inner]
	}
	fmt.Println(inputmatrix)
}
func main() {
	inputmatrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	for i := 0; i < len(inputmatrix)/2; i++ {
		rotatematrix(inputmatrix, len(inputmatrix), i)
	}
}
