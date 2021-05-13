package palindromic_substrings

/*
	
	Given a string s, return the number of palindromic substrings in it.

	Input: s = "abc"
	Explanation: Three palindromic strings: "a", "b", "c".

	Input: s = "aaa"
	Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

	Approach -> If any string is supposed to be a palindrome, then the string[start+1:end] should also be a palindrome
 */

func countSubstrings(s string) int {
	sSplit := strings.Split(s, "")
	prevData := make([][]bool, len(sSplit))
	var trueCount int
	for i := 0; i < len(prevData); i++ {
		prevData[i] = make([]bool, len(sSplit))
		prevData[i][i] = true
		trueCount++
	}
	for i := 1; i < len(sSplit); i++ {
		for j := 0; j < len(sSplit)-i; j++ {
			startIndex := j
			endIndex := j + i
			if startIndex+1 >= endIndex {
				if sSplit[startIndex] == sSplit[endIndex] {
					prevData[startIndex][endIndex] = true
					trueCount++
				}
				continue
			}
			if prevData[startIndex+1][endIndex-1] == true && sSplit[startIndex] == sSplit[endIndex] {
				prevData[startIndex][endIndex] = true
				trueCount++
			}
		}
	}
	return trueCount
}
