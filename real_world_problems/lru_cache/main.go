package main

type node struct {
	left  *node
	key   int
	value int
	right *node
}

type LRUCache struct {
	head        *node
	last        *node
	keyNodeMap  map[int]*node
	filledCount int
	capacity    int
}

func InitialiseCacheObject(capacity int) LRUCache {
	return LRUCache{
		head:        nil,
		last:        nil,
		keyNodeMap:  make(map[int]*node,capacity),
		filledCount: 0,
		capacity:    capacity,
	}
}

func createNewNode(key int, value int) *node {
	return &node{
		key:   key,
		value: value,
	}
}

func (cacheObject *LRUCache) moveNodeToLast(nodeAddress *node) {
	if cacheObject.last == nodeAddress {
		return
	}
	if nodeAddress == cacheObject.head {
		if nodeAddress.right != nil {
			cacheObject.head = nodeAddress.right
			nodeAddress.right.left = nil
		}
	} else {
		if nodeAddress.right != nil {
			nodeAddress.left.right = nodeAddress.right
			nodeAddress.right.left = nodeAddress.left
		}
	}
	cacheObject.last.right = nodeAddress
	nodeAddress.left = cacheObject.last
	nodeAddress.right = nil
	cacheObject.last = nodeAddress
}

func (cacheObject *LRUCache) Get(key int) int {
	nodeAddress, nodeExists := cacheObject.keyNodeMap[key]
	if !nodeExists {
		return -1
	}
	cacheObject.moveNodeToLast(nodeAddress)
	return nodeAddress.value
}

func (cacheObject *LRUCache) Put(key int, value int) {

	nodeAddress, nodeExists := cacheObject.keyNodeMap[key]
	if nodeExists {
		cacheObject.moveNodeToLast(nodeAddress)
		nodeAddress.value = value
		return
	}
	if cacheObject.capacity == cacheObject.filledCount {
		headValue := cacheObject.head.key
		if cacheObject.head.right != nil {
			cacheObject.head.right.left = nil
		}
		cacheObject.head = cacheObject.head.right
		newNode := createNewNode(key, value)
		newNode.left = cacheObject.last
		cacheObject.last.right = newNode
		cacheObject.last = newNode
		delete(cacheObject.keyNodeMap, headValue)
		cacheObject.keyNodeMap[key] = newNode
		return
	}
	newNode := createNewNode(key, value)
	if cacheObject.head == nil {
		cacheObject.head = newNode
		cacheObject.last = newNode
	} else {
		newNode.left = cacheObject.last
		cacheObject.last.right = newNode
		cacheObject.last = newNode
	}
	cacheObject.keyNodeMap[key] = newNode
	cacheObject.filledCount = cacheObject.filledCount + 1
}


// Usage ----

func main() {
	cacheObject := InitialiseCacheObject(10)
	cacheObject.Put(1,10)
	cacheObject.Put(2, 30)
	cacheObject.Get(3)
	cacheObject.Get(1)
	cacheObject.Put(5,50)
}