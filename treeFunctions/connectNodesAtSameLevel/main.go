package connectNodesAtSameLevel

import (
	"go-dsa/treeformation"
	"fmt"
)

func main() {
	treeNodes := []int64{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5}
	headNode := treeformation.CreateTree(treeNodes)

}
