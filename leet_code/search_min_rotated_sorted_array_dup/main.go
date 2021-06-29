package main

func linearSearch(nums []int, startIndex, endIndex, mid int) (min int) {
	for i := startIndex; i < endIndex; i++ {
		if nums[i] < mid {
			return nums[i]
		}
	}
	return mid
}
func findMin(nums []int) int {
	startIndex := 0
	endIndex := len(nums) - 1
	mid := (startIndex + endIndex) / 2
	for {
		if mid == startIndex || mid == endIndex {
			if nums[startIndex] > nums[endIndex] {
				return nums[endIndex]
			}
			return nums[startIndex]
		}
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if nums[mid] > nums[startIndex] && nums[mid] < nums[endIndex-1] {
			return nums[startIndex]
		}
		if nums[mid] > nums[endIndex] && nums[mid] < nums[startIndex] {
			startIndex = mid + 1
		} else if nums[mid] < nums[endIndex] && nums[mid] > nums[startIndex] {
			endIndex = mid
		} else {
			/*
				In case there is a tie deciding which direction to move, we don't have any other choice
				than iterating through the entire slice. (Worst case will be O(n)
			*/
			rightMin := linearSearch(nums, mid+1, endIndex+1, nums[mid])
			leftMin := linearSearch(nums, startIndex, mid, nums[mid])
			if rightMin > leftMin {
				return leftMin
			}
			return rightMin
		}
		mid = (startIndex + endIndex) / 2
	}
	return 0
}
