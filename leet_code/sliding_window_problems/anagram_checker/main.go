package main

import "strings"

/*
	Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

	Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

	The order of output does not matter.

	Example 1:

	Input:
	s: "cbaebabacd" p: "abc"

	Output:
	[0, 6]

	Explanation:
	The substring with start index = 0 is "cba", which is an anagram of "abc".
	The substring with start index = 6 is "bac", which is an anagram of "abc".
	Example 2:

	Input:
	s: "abab" p: "ab"

	Output:
	[0, 1, 2]

	Explanation:
	The substring with start index = 0 is "ab", which is an anagram of "ab".
	The substring with start index = 1 is "ba", which is an anagram of "ab".
	The substring with start index = 2 is "ab", which is an anagram of "ab".

*/
func findAnagrams(mainString string, request string) []int {
	if len(mainString) < len(request) {
		return []int{}
	}
	requestMap := make(map[string]int)
	var requestlen int
	var requestChar string
	for requestlen, requestChar = range strings.Split(request, "") {
		requestMap[requestChar]++
	}
	sub := mainString[0 : requestlen+1]
	subMap := make(map[string]int)
	for _, subChar := range strings.Split(sub, "") {
		subMap[subChar]++
	}
	matchingData := make(map[string]int)
	var missingCount int
	for key, requestKeyCount := range requestMap {
		if subMap[key] < requestKeyCount {
			missingCount += requestKeyCount - subMap[key]
			matchingData[key] = subMap[key]
		}
		matchingData[key] = subMap[key]
	}
	finalList := []int{}
	if missingCount == 0 {
		finalList = append(finalList, 0)
	}
	for i := 1; i < len(mainString)-requestlen; i++ {
		newEntry := string(mainString[i+requestlen])
		oldEntry := string(mainString[i-1])
		if requestMap[oldEntry] > 0 {
			if matchingData[oldEntry] <= requestMap[oldEntry] {
				missingCount++
			}
			matchingData[oldEntry]--
		}
		if requestMap[newEntry] > 0 {
			if matchingData[newEntry] < requestMap[newEntry] {
				missingCount--
			}
			matchingData[newEntry]++
		}
		if missingCount == 0 {
			finalList = append(finalList, i)
		}

	}
	return finalList
}

func main() {
	// Usage :
	findAnagrams("cbaebabacd", "abc")
}
