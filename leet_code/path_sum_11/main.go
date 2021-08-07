package path_sum_11

import "go-dsa/treeformation"

/*

	Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where each path's sum equals targetSum.

	A leaf is a node with no children.

	Example 1:

	Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
	Output: [[5,4,11,2],[5,8,4,5]]

*/
func findPathSum(root *treeformation.Node, actualSum, targetSum int64, path []int64) [][]int64 {
	if root.Left == nil && root.Right == nil {
		if actualSum+root.Data == targetSum {
			finalPath := append(path, root.Data)
			return [][]int64{finalPath}
		}
	}
	var leftPaths, rightPaths [][]int64
	if root.Left != nil {
		lenCurPath := len(path)
		tmp := make([]int64, lenCurPath, lenCurPath+1)
		_ = copy(tmp, path)
		leftPaths = findPathSum(root.Left, actualSum+root.Data, targetSum, append(tmp, root.Data))
	}

	if root.Right != nil {
		lenCurPath := len(path)
		tmp := make([]int64, lenCurPath, lenCurPath+1)
		_ = copy(tmp, path)
		rightPaths = findPathSum(root.Right, actualSum+root.Data, targetSum, append(tmp, root.Data))
	}
	return append(leftPaths, rightPaths...)
}

func pathSum(root *treeformation.Node, targetSum int64) [][]int64 {
	if root == nil {
		return [][]int64{}
	}
	return findPathSum(root, 0, targetSum, []int64{})
}
