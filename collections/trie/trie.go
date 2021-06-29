package trie

import (
	"fmt"
	"go-dsa/collections"
)

type trie struct {
	children map[rune]*trie
	edgeData []interface{}
}

func (t *trie) InsertNode(s string, dataToInsertAtEnd interface{}) {
	root := t
	for index, char := range s {
		if _, ok := root.children[char]; !ok {
			newNode := &trie{
				children: map[rune]*trie{},
				edgeData: nil,
			}
			root.children[char] = newNode
			if index == len(s)-1 {
				root.children[char].edgeData = append(root.children[char].edgeData, dataToInsertAtEnd)
				return
			}
			root = root.children[char]
		} else {
			if index == len(s)-1 {
				root.children[char].edgeData = append(root.children[char].edgeData, dataToInsertAtEnd)
				return
			}
			root = root.children[char]
		}
	}
}

func (t *trie) SearchNode(s string) bool {
	current := t
	for index, char := range s {
		if _, ok := current.children[char]; !ok {
			return false
		}
		if index == len(s)-1 {
			if current.children[char].edgeData == nil {
				return false
			}
			return true
		}
		current = current.children[char]
	}
	return false
}

func (t *trie) Print() {
	fmt.Println(getSubStrings(t, ""))
}

func getSubStrings(startNode *trie, str string) (result []collections.Metadata) {
	for char := range startNode.children {
		result = append(result, getSubStrings(startNode.children[char], str+string(char))...)
		if startNode.children[char].edgeData != nil {
			result = append(result, collections.Metadata{
				SuffixString: str + string(char),
				DataStored:   startNode.children[char].edgeData,
			})
		}
	}
	return result
}

func (t *trie) GetDataStartingWith(startsWith string) (result []collections.Metadata) {
	current := t
	for index, char := range startsWith {
		if _, ok := current.children[char]; !ok {
			return result
		}
		if current.children[char].edgeData != nil && index != len(startsWith)-1 {
			result = append(result, collections.Metadata{
				SuffixString: startsWith[index+1:],
				DataStored:   current.children[char].edgeData,
			})
		}
		if index == len(startsWith)-1 {
			result = append(result, getSubStrings(current.children[char], "")...)
			if current.children[char].edgeData != nil {
				result = append(result, collections.Metadata{
					SuffixString: "",
					DataStored:   current.children[char].edgeData,
				})
			}
			return result
		}
		current = current.children[char]
	}
	return result
}

func New() collections.Trie {
	return &trie{
		children: map[rune]*trie{},
	}
}
