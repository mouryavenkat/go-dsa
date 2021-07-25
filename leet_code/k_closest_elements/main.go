package main

import (
	"fmt"
	"go-dsa/collections/searches"
	"go-dsa/helpers"
)

func getKClosest(elements []int, k, x int) []int {
	obj := searches.NewBinarySearch()
	index := obj.BinarySearchClosest(elements, x)
	count := 0
	leftIndex := index
	rightIndex := index
	if rightIndex == len(elements) - 1 {
		count = 1
		rightIndex += 1
	}
	fmt.Println(index)
	for count < k {
		if rightIndex > len(elements)-1 || (leftIndex > 0 && helpers.Abs(elements[leftIndex]-x) <= helpers.Abs(elements[rightIndex]-x)) {
			leftIndex = leftIndex - 1
		} else if leftIndex <= 0 || (rightIndex < len(elements)-1 && helpers.Abs(elements[leftIndex]-x) >= helpers.Abs(elements[rightIndex]-x)) {
			rightIndex = rightIndex + 1
		}
		count = count + 1
	}
	if rightIndex > len(elements) - 1 {
		return elements[leftIndex:]
	}
	fmt.Println(leftIndex, rightIndex)
	return elements[leftIndex:rightIndex]
}

func main() {
	fmt.Println(getKClosest([]int{1, 2, 3, 7, 8, 9}, 1, 7))
}
