package main

import (
	"math"
	"strings"
)

/*
	Given a string s and a character c that occurs in s, return an array of integers answer where answer.length == s.length and answer[i] is the distance from index i to the closest occurrence of character c in s.

	The distance between two indices i and j is abs(i - j), where abs is the absolute value function.

	Example 1:

	Input: s = "loveleetcode", c = "e"
	Output: [3,2,1,0,1,0,0,1,2,2,1,0]
	Explanation: The character 'e' appears at indices 3, 5, 6, and 11 (0-indexed).
	The closest occurrence of 'e' for index 0 is at index 3, so the distance is abs(0 - 3) = 3.
	The closest occurrence of 'e' for index 1 is at index 3, so the distance is abs(1 - 3) = 3.
	For index 4, there is a tie between the 'e' at index 3 and the 'e' at index 5, but the distance is still the same: abs(4 - 3) == abs(4 - 5) = 1.
	The closest occurrence of 'e' for index 8 is at index 6, so the distance is abs(8 - 6) = 2.
*/

func shortestToChar(s string, c byte) []int {
	charSplit := strings.Split(s, "")
	leftClosest := make([]float64, len(charSplit))
	rightClosest := make([]float64, len(charSplit))
	finalArr := make([]int, len(charSplit))
	var closestLeftIndex = math.Inf(1)
	var closestRightIndex = math.Inf(1)
	for i := range charSplit {
		if charSplit[i] == string(c) {
			closestLeftIndex = float64(i)
		}
		leftClosest[i] = closestLeftIndex
		if charSplit[len(charSplit)-i-1] == string(c) {
			closestRightIndex = float64(len(charSplit) - i - 1)
		}
		rightClosest[len(charSplit)-i-1] = closestRightIndex
	}
	for i := range charSplit {
		if math.Abs(leftClosest[i]-float64(i)) > math.Abs(rightClosest[i]-float64(i)) {
			finalArr[i] = int(math.Abs(rightClosest[i] - float64(i)))
		} else {
			finalArr[i] = int(math.Abs(leftClosest[i] - float64(i)))
		}
	}
	return finalArr
}

func main() {
	shortestToChar("loveleetcode", 'e')
}
