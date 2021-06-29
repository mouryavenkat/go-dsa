package main

import "fmt"

/*
	Given an integer array nums, return the length of the longest strictly increasing subsequence.

	A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].

	Input: nums = [10,9,2,5,3,7,101,18]
	Output: 4
	Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
*/
func lengthOfLIS(nums []int) int {
	longestSequence := make([]int, len(nums))
	var max int
	for i := len(nums) - 1; i >= 0; i-- {
		var maxCount int
		for j := i; j < len(nums); j++ {
			if longestSequence[j] > maxCount && nums[i] < nums[j] {
				maxCount = longestSequence[j]
			}
		}
		longestSequence[i] = maxCount + 1
		if max < maxCount+1 {
			max = maxCount + 1
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
