package main

import "go-dsa/treeformation"

/*
		You need to find the largest value in each row of a binary tree.

Example:
Input:

          1
         / \
        3   2
       / \   \
      5   3   9

Output: [1, 3, 9]

*/
func largestValues(root *treeformation.Node) []int {
	if root == nil {
		return []int{}
	}
	cNodes := []*treeformation.Node{root}
	cNodesCount := 1
	bestArr := make([]int, 0)
	for cNodesCount > 0 {
		nextLevelNodes := make([]*treeformation.Node, 0)
		var nextLevelNodesCount int
		bestInLevel := cNodes[0].Data
		for i := 0; i < cNodesCount; i++ {
			if cNodes[i].Left != nil {
				nextLevelNodes = append(nextLevelNodes, cNodes[i].Left)
				nextLevelNodesCount++
			}
			if cNodes[i].Right != nil {
				nextLevelNodes = append(nextLevelNodes, cNodes[i].Right)
				nextLevelNodesCount++
			}
			if cNodes[i].Data > bestInLevel {
				bestInLevel = cNodes[i].Data
			}

		}
		cNodes = nextLevelNodes
		cNodesCount = nextLevelNodesCount
		bestArr = append(bestArr, int(bestInLevel))
	}
	return bestArr
}
