package main

import "sort"

/*
	Given an integer array nums of size n, return the minimum number of moves required to make all array elements equal.

	In one move, you can increment or decrement an element of the array by 1

	Input: nums = [1,2,3]
	Output: 2
	Explanation:
	Only two moves are needed (remember each move increments or decrements one element):
	[1,2,3]  =>  [2,2,3]  =>  [2,2,2]

	Input: nums = [1,10,2,9]
	Output: 16
*/
func minMoves2(nums []int) int {
	sort.Ints(nums)
	var moves int
	if len(nums)%2 != 0 {
		midIndex := (len(nums) - 1) / 2
		for i := 0; i < midIndex; i++ {
			moves += nums[midIndex] - nums[i]
		}
		for i := midIndex + 1; i < len(nums); i++ {
			moves += nums[i] - nums[midIndex]
		}
	} else {
		midElement := (nums[len(nums)/2] + nums[(len(nums)-2)/2]) / 2
		var i int
		for nums[i] <= midElement {
			moves += midElement - nums[i]
			i++
		}
		for i < len(nums) {
			moves += nums[i] - midElement
			i++
		}
	}
	return moves
}
