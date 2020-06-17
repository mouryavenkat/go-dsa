package main

import "sort"

/*
	Given a collection of intervals, merge all overlapping intervals.

	Example 1:

	Input: [[1,3],[2,6],[8,10],[15,18]]
	Output: [[1,6],[8,10],[15,18]]
	Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
	Example 2:

	Input: [[1,4],[4,5]]
	Output: [[1,5]]
	Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	i := 0
	intLen := len(intervals)
	newIntervals := [][]int{}
	for i < intLen {
		temp := i
		bestRight := intervals[i][1]
		for i < intLen-1 && (intervals[i+1][0] >= intervals[temp][0] && intervals[i+1][0] <= bestRight) {
			if bestRight < max(intervals[i][1], intervals[i+1][1]) {
				bestRight = max(intervals[i][1], intervals[i+1][1])
			}
			i++
		}

		newIntervals = append(newIntervals, []int{intervals[temp][0], bestRight})
		i++
	}
	return newIntervals
}
