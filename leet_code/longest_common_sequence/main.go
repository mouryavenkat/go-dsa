package main

/*
	Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

	You must write an algorithm that runs in O(n) time.

	Example 1:

	Input: nums = [100,4,200,1,3,2]
	Output: 4
	Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
	Example 2:

	Input: nums = [0,3,7,2,5,8,4,6,0,1]
	Output: 9
*/
func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = false
	}
	var maxCount int
	for _, num := range nums {
		isVisited, _ := numsMap[num]
		if isVisited {
			continue
		}
		numsMap[num] = true
		surroundings := []int{num}
		count := 1
		for len(surroundings) != 0 {
			poppedElement := surroundings[0]
			if isVisited, ok := numsMap[poppedElement+1]; ok && !isVisited {
				numsMap[poppedElement+1] = true
				surroundings = append(surroundings, poppedElement+1)
				count = count + 1
			}
			if isVisited, ok := numsMap[poppedElement-1]; ok && !isVisited {
				numsMap[poppedElement-1] = true
				surroundings = append(surroundings, poppedElement-1)
				count = count + 1
			}
			if len(surroundings) > 0 {
				surroundings = surroundings[1:]
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func main() {
	longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
}
