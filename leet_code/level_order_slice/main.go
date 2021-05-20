package main

import "go-dsa/treeformation"

/*
	Input: root = [3,9,20,null,null,15,7]
	Output: [[3],[9,20],[15,7]]
 */
func getOrderedElements(root *treeformation.Node, finalSlice [][]int64, index int) [][]int64 {
	if root == nil {
		return finalSlice
	}
	if len(finalSlice) == index {
		finalSlice = append(finalSlice, []int64{root.Data})

	} else {
		finalSlice[index] = append(finalSlice[index], root.Data)
	}
	index = index + 1
	finalSlice = getOrderedElements(root.Left, finalSlice, index)
	finalSlice = getOrderedElements(root.Right, finalSlice, index)
	return finalSlice
}

func levelOrder(root *treeformation.Node) [][]int64 {
	return getOrderedElements(root, [][]int64{}, 0)
}
