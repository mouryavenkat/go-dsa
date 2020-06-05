package main

import (
	"go-dsa/treeformation"
	"math"
)

/*
		Given a binary tree, determine if it is a valid binary search tree (BST).

		Assume a BST is defined as follows:

		The left subtree of a node contains only nodes with keys less than the node's key.
		The right subtree of a node contains only nodes with keys greater than the node's key.
		Both the left and right subtrees must also be binary search trees.


		Example 1:

			2
		   / \
		  1   3

		Input: [2,1,3]
		Output: true
		Example 2:

			5
		   / \
		  1   4
			 / \
			3   6

		Input: [5,1,4,null,null,3,6]
		Output: false
		Explanation: The root node's value is 5 but its right child's value is 4.
 */

func checkBST(root *treeformation.Node, min, max int64) bool {
	if root == nil {
		return true
	}
	if root.Data <= min || root.Data>= max {
		return false
	}
	return checkBST(root.Left, min, root.Data) && checkBST(root.Right, root.Data, max)
}
func isValidBST(root *treeformation.Node) bool {
	return checkBST(root, -math.MaxInt64, math.MaxInt64)
}
