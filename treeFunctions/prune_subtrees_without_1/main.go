package main

import "go-dsa/treeformation"

/*
	Given the root of a binary tree, return the same tree where every subtree (of the given tree) not containing a 1 has been removed.

	A subtree of a node node is node plus every node that is a descendant of node.

	Example 1:

	Input: root = [1,null,0,0,1]
	Output: [1,null,0,null,1]
	Explanation:
	Only the red nodes satisfy the property "every subtree not containing a 1".
	The diagram on the right represents the answer.

	Example 2:

	Input: root = [1,0,1,0,0,0,1]
	Output: [1,null,1,null,1]

*/
func prune(root *treeformation.Node) bool {
	if root == nil {
		return false
	}
	left := prune(root.Left)
	if !left {
		root.Left = nil
	}
	right := prune(root.Right)
	if !right {
		root.Right = nil
	}
	return left || right || root.Data == 1
}

func pruneTree(root *treeformation.Node) *treeformation.Node {
	if root == nil {
		return nil
	}
	prune(root)
	if root.Left == nil && root.Right == nil && root.Data == 0 {
		return nil
	}
	return root
}
