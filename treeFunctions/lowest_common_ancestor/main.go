package main

import (
	"fmt"
	"go-dsa/treeformation"
)

/*
	Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

	According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

	Example 1:

	Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
	Output: 6
	Explanation: The LCA of nodes 2 and 8 is 6.
 */

func lowestCommonAncestor(root, p, q *treeformation.Node) *treeformation.Node {
	if root == nil {
		fmt.Println("Root is nil")
		return root
	}
	if p.Data > q.Data {
		p, q = q, p
	}
	if root.Data >= p.Data && root.Data <= q.Data {
		return root
	} else if root.Data > p.Data {
		root = lowestCommonAncestor(root.Left, p, q)
	} else if root.Data < q.Data {
		root = lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
