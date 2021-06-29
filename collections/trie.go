package collections

type Trie interface {
	InsertNode(s string, dataToInsertAtEnd interface{})
	SearchNode(s string) bool
	Print()
	GetDataStartingWith(startsWith string) []Metadata
}

type Metadata struct {
	SuffixString string
	DataStored   interface{}
}
