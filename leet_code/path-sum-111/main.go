package main

import "go-dsa/treeformation"

/*
	You are given a binary tree in which each node contains an integer value.

	Find the number of paths that sum to a given value.

	The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

	The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

	Example:

	root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

		  10
		 /  \
		5   -3
	   / \    \
	  3   2   11
	 / \   \
	3  -2   1

	Return 3. The paths that sum to 8 are:

	1.  5 -> 3
	2.  5 -> 2 -> 1
	3. -3 -> 11
*/

func calculatePathCount(root *treeformation.Node, sum int64, pathSlice []int64) int {
	if root == nil {
		return 0
	}
	var count int
	var tempSlice []int64
	for index := range pathSlice {
		tempSlice = append(tempSlice, pathSlice[index]+root.Data)
		if tempSlice[index] == sum {
			count++
		}
	}
	if root.Data == sum {
		count++
	}
	tempSlice = append(tempSlice, root.Data)
	return count + calculatePathCount(root.Left, sum, tempSlice) + calculatePathCount(root.Right, sum, tempSlice)
}

func pathSum(root *treeformation.Node, sum int64) int {
	return calculatePathCount(root, sum, []int64{})
}
