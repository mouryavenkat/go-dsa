package main

import (
	"go-dsa/treeformation"
)

/*
	Given a binary tree where node values are digits from 1 to 9. A path in the binary tree is said to be pseudo-palindromic if at least one permutation of the node values in the path is a palindrome.

	Return the number of pseudo-palindromic paths going from the root node to leaf nodes.

	Example 1:

	Input: root = [2,3,1,3,1,null,1]
	Output: 2
	Explanation: The figure above represents the given binary tree.
	There are three paths going from the root node to leaf nodes: the red path [2,3,3], the green path [2,1,1], and the path [2,3,1].
	Among these paths only red path and green path are pseudo-palindromic paths since the red path [2,3,3] can be rearranged in [3,2,3] (
	palindrome) and the green path [2,1,1] can be rearranged in [1,2,1] (palindrome).
*/

func recursive(root *treeformation.Node, path []int64) (finalCount int) {
	if root == nil {
		return
	}
	path = append(path, root.Data)
	finalCount += recursive(root.Left, path)
	finalCount += recursive(root.Right, path)
	if root.Left == nil && root.Right == nil {
		oddCount := 0
		pathCount := map[int64]int{}
		for _, element := range path {
			count, exists := pathCount[element]
			if !exists {
				pathCount[element] = 1
				oddCount++
				continue
			}
			if count%2 != 0 {
				oddCount--
			} else {
				oddCount++
			}
			pathCount[element]++
		}
		if oddCount <= 1 {
			finalCount++
		}
	}
	return finalCount
}
func pseudoPalindromicPaths(root *treeformation.Node) int {
	return recursive(root, []int64{})
}
