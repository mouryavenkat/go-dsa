package main

import "strings"

/*
	Given two strings A and B of lowercase letters, return true if you can swap two letters in A so the result is equal to B, otherwise, return false.

	Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at A[i] and A[j]. For example, swapping at indices 0 and 2 in "abcd" results in "cbad".

	Example 1:

	Input: A = "ab", B = "ba"
	Output: true
	Explanation: You can swap A[0] = 'a' and A[1] = 'b' to get "ba", which is equal to B.
	Example 2:

	Input: A = "ab", B = "ab"
	Output: false
	Explanation: The only letters you can swap are A[0] = 'a' and A[1] = 'b', which results in "ba" != B.
*/
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	aMap := make(map[string]int, len(A))
	aSplit := strings.Split(A, "")
	bSplit := strings.Split(B, "")
	aDiff := make([]string, 2, 2)
	bDiff := make([]string, 2 ,2)
	mismatchCount := 0
	for i := 0; i < len(aSplit); i++ {
		aMap[aSplit[i]]++
		if mismatchCount > 2 {
			return false
		}
		if aSplit[i] != bSplit[i] {
			aDiff[mismatchCount] = aSplit[i]
			bDiff[mismatchCount] = bSplit[i]
			mismatchCount++
		}
	}
	// fmt.Println(aMap, bMap, aDiff, bDiff)
	if mismatchCount == 1 {
		return false
	}
	isBuddy := true
	if len(aDiff) == 0 {
		isBuddy = false
		for _, value := range aMap {
			if value >= 2 {
				return true
			}
		}
		return isBuddy
	}
	if aDiff[0] != bDiff[1] || aDiff[1] != bDiff[0] {
		return false
	}

	return isBuddy
}
func main() {
	buddyStrings("aaab", "aaba")
}
