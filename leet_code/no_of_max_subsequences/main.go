package no_of_max_subsequences

import "strings"

func next(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	ans := -1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] <= target {
			start = mid + 1
		} else {
			ans = mid
			end = mid - 1
		}
	}
	return ans
}

/*

	Given a string s and an array of strings words, return the number of words[i] that is a subsequence of s.

	A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

	For example, "ace" is a subsequence of "abcde".


	Example 1:

	Input: s = "abcde", words = ["a","bb","acd","ace"]
	Output: 3
	Explanation: There are three strings in words that are a subsequence of s: "a", "acd", "ace".
	Example 2:

	Input: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
	Output: 2
*/

func numMatchingSubsequences(s string, words []string) int {
	actualSMap := map[string][]int{}
	for index, char := range strings.Split(s, "") {
		actualSMap[char] = append(actualSMap[char], index)
	}
	validCount := 0
	for _, word := range words {
		sMapCopy := map[string][]int{}
		for char, data := range actualSMap {
			sMapCopy[char] = data
		}
		lastIndex := -1
		wSplit := strings.Split(word, "")
		isValidWord := true

		for _, char := range wSplit {
			if data, ok := sMapCopy[char]; !ok || len(data) == 0 {
				isValidWord = false
				break
			}
			if sMapCopy[char][0] <= lastIndex {
				target := next(sMapCopy[char], lastIndex)
				if target == -1 {
					isValidWord = false
					break
				}
				sMapCopy[char] = sMapCopy[char][target:]
			}
			lastIndex = sMapCopy[char][0]
			sMapCopy[char] = sMapCopy[char][1:]
		}
		if isValidWord {
			validCount++
		}
	}
	return validCount
}
