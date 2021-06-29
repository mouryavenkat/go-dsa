package main

import (
	heap2 "go-dsa/collections/heap"
)

func topKFrequent(nums []int, k int) []int {
	h := heap2.NewHeap()
	numsMap := map[int]int{}
	for _, num := range nums {
		numsMap[num]++
	}
	for key, value := range numsMap {
		h.Insert(&heap2.Node{
			Val:  key,
			Data: value,
		})
	}
	returnData := make([]int, k)
	for i := 0; i < k; i++ {
		returnData[i] = h.Delete().Val

	}
	return returnData
}

func main() {
	/*
		Given a non-empty array of integers, return the k most frequent elements.

		Example 1:

		Input: nums = [1,1,1,2,2,3], k = 2
		Output: [1,2]
		Example 2:

		Input: nums = [1], k = 1
		Output: [1]
		Note:

		You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
		Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
		It's guaranteed that the answer is unique, in other words the set of the top k frequent elements is unique.
		You can return the answer in any order.
	*/
}
