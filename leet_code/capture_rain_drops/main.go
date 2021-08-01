package capture_rain_drops

/*
	Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

	Example 1:

	Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
	Output: 6
	Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
	Example 2:

	Input: height = [4,2,0,3,2,5]
	Output: 9

*/
func getMax(height []int, start, end int) int {
	if start >= len(height) {
		return len(height) - 1
	}
	max := start
	for i := start; i <= end; i++ {
		if height[i] > height[max] {
			max = i
		}
	}
	return max
}

func trap(height []int) int {
	var stack []int
	nextBestIndex := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		if height[stack[len(stack)-1]] <= height[i] {
			for len(stack) > 0 && height[stack[len(stack)-1]] <= height[i] {
				nextBestIndex[stack[len(stack)-1]] = i
				stack = stack[0 : len(stack)-1]
			}
		}
		stack = append(stack, i)
	}
	lastIndex := len(height) - 1
	for len(stack) > 0 {
		nextBestIndex[stack[len(stack)-1]] = getMax(height, stack[len(stack)-1]+1, lastIndex)
		lastIndex = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
	}
	sum := 0
	i := 0
	for i < len(height) {
		nextBest := nextBestIndex[i]
		if nextBest == i {
			return sum
		}
		modifiedHeight := height[i]
		if height[i] > height[nextBest] {
			modifiedHeight = height[nextBest]
		}
		for j := i + 1; j < nextBest; j++ {
			sum = sum + modifiedHeight - height[j]
		}
		i = nextBest
	}
	return sum
}
