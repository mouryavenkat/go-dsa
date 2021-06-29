package main

import (
	"fmt"
	"go-dsa/collections/trie"
	"go-dsa/string_funtions"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	words := []string{"bbbbbabbabbbb", "baabbaa", "bbab", "bbbabbaaab", "abababbbbbab", "abb", "baaaabbb", "babbaaaba", "aab", "aaaab", "baabbbbabbaaaba", "baaab", "abbbab", "abaabbbabbabba", "aa", "aabbbaabba", "aaabbbbbaaabbbb", "bbaaaaba", "ababaaa", "aaaaa", "aaaaabbbbaaaba", "abbabbbaabbaabbb", "bbaba", "aaaaabbbabbbbaaaab", "abbbaa", "bbbabbaaa", "bbbaaabaabbbaaaaabaa", "aaaabbabb", "ababbababbbab", "aaaaababaababbbabaaa", "ba", "bbbbababbbabab", "baaaba", "aababbaaabbb", "aabbaaabbabaaababaab", "abbbb", "babaabaaababb", "bbbbabaaaab", "babbbbb", "babaaba", "aaba", "abababba", "a", "bb", "abaaab", "babbabaababbabaaba", "aaaaaababbbabaaabaa", "baabaaabb", "b", "bbaaaabbb", "abaaaaabaabbbaa", "ab", "bababaaaba", "aabababb", "ababaabbaababba", "bbb", "ababbaabababbbbbabb", "bbbbb", "abbbbaabaaaabb", "baba", "bbaabbabaaababaabbaa", "bbaabaabbabbaab", "bbbaabbab", "babbbbbaaaaabaa", "abbbbbbabbbabb", "abaa", "bbbbaababaab", "abaaababa", "aaaababaaababbaaba", "bbabbbabbbbbbaab", "abbabbabaabbabbbba", "abbbbbaabbbaaabaaaa", "bbaabababb", "aaabababaabbaaaaaaab", "ababbaabaaababb", "abbbbabaaabbaaabbab", "aababababbabaaa", "baabbaabbbaaaaaa", "bbbbbbbabbabbbbbb", "bbbabbabbabbabaabba", "babbbbabaaaabbabaab", "baabaabaabababaaabba", "bbaaaabbbbabbbaaaa", "aaaaabaabaa", "bbabaaabbbabaa", "baaabaaaaaab", "ababbbbbbbabaaaba", "abbbabaababbbbbaaa", "baaaaaabab", "aabbabba", "baaabbaabbbbaba", "aaaaabba", "babaaabbba", "bbbbab", "bbbbaabbaabab", "baa", "baababaa", "abbbbb", "babbaa", "abbbabbaa"}
	trieObject := trie.New()
	for index, input := range words {
		trieObject.InsertNode(reverse(input), index)
	}

	var result [][]int
	var palindromeStringIndexes []int
	for index, input := range words {
		if input == "" {
			continue
		}
		if string_funtions.CheckIsPalindrome(input) {
			palindromeStringIndexes = append(palindromeStringIndexes, index)
		}
	}
	for index, input := range words {
		if input == "" {
			for _, palindromeIndex := range palindromeStringIndexes {
				result = append(result, []int{index, palindromeIndex})
				result = append(result, []int{palindromeIndex, index})
			}
		}
		probableStrings := trieObject.GetDataStartingWith(input)
		for _, probablePalindromesData := range probableStrings {
			if string_funtions.CheckIsPalindrome(probablePalindromesData.SuffixString) {
				dataStored := probablePalindromesData.DataStored.([]interface{})
				for _, data := range dataStored {
					if data.(int) != index {
						result = append(result, []int{index, data.(int)})
					}
				}
			}
		}
	}
	fmt.Println(result)
}
