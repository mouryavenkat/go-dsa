package main

import "strings"

/*
	Given a string s, a k duplicate removal consists of choosing k adjacent and equal letters from s and removing them causing the left and the right side of the deleted substring to concatenate together.

	We repeatedly make k duplicate removals on s until we no longer can.

	Return the final string after all such duplicate removals have been made.

	It is guaranteed that the answer is unique.

 */
func removeDuplicates(s string, k int) string {
	sSplit := strings.Split(s, "")
	charMap := map[string]bool{}
	for _, char := range sSplit {
		charMap[char] = true
	}
	contains := true
	for contains {
		contains = false
		for char := range charMap {
			repChar := strings.Repeat(char, k)
			if strings.Contains(s, repChar) {
				contains = true
			}
			s = strings.Replace(s, repChar, "", -1)
		}
	}
	return s
}
