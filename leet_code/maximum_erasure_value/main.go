package main

/*
	You are given an array of positive integers nums and want to erase a subarray containing unique elements. The score you get by erasing the subarray is equal to the sum of its elements.

	Return the maximum score you can get by erasing exactly one subarray.

	An array b is called to be a subarray of a if it forms a contiguous subsequence of a, that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).

	Example 1:

	Input: nums = [4,2,4,5,6]
	Output: 17
	Explanation: The optimal subarray here is [2,4,5,6].
	Example 2:

	Input: nums = [5,2,1,2,5,2,1,2,5]
	Output: 8
	Explanation: The optimal subarray here is [5,2,1] or [1,2,5].
 */

func maximumUniqueSubarray(nums []int) int {
	cumSum := make([]int, len(nums))
	var sum int
	for index, num := range nums {
		cumSum[index] = sum
		sum = sum + num
	}
	var bestSum int
	numsMap := map[int]int{}
	startIndex := 0
	for index, num := range nums {
		lastIndex, ok := numsMap[num]
		if ok && lastIndex >= startIndex {
			sum := cumSum[index] - cumSum[startIndex]
			if sum > bestSum {
				bestSum = sum
			}
			startIndex = lastIndex + 1
		}
		numsMap[num] = index
	}
	if startIndex <= len(nums)-1 {
		sum = cumSum[len(nums)-1] - cumSum[startIndex] + nums[len(nums)-1]
		if sum > bestSum {
			bestSum = sum
		}
	}
	return bestSum
}
