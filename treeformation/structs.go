package treeformation

type Node struct {
	Data  int64
	Left  *Node
	Right *Node
}

type NodeWihNextRight struct {
	Data      int64
	Left      *NodeWihNextRight
	Right     *NodeWihNextRight
	NextRight *NodeWihNextRight
}
