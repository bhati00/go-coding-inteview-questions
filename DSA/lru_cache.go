package languagebasics

/*
LRU Cache Implementation

Core Idea:
- Most recently used items are kept at the front (near head)
- Least recently used items move toward the back (near tail)
- When capacity is full, remove from the back (least recently used)

Data Structures:
- HashMap: O(1) lookup to find nodes by key
- Doubly Linked List: O(1) deletion and insertion
  - head (dummy) -> [most recent] -> ... -> [least recent] -> tail (dummy)
*/

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LruCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func NewLruCache(capacity int) *LruCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &LruCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (lru *LruCache) Get(key int) int {
	if node, exists := lru.cache[key]; exists {
		lru.moveToFront(node)
		return node.val
	}
	return -1
}

// Put adds or updates a key-value pair
func (lru *LruCache) Put(key int, val int) {
	// If key already exists, update it and move to front
	if node, exists := lru.cache[key]; exists {
		node.val = val
		lru.moveToFront(node)
		return
	}

	// If at capacity, remove least recently used (from back)
	if len(lru.cache) >= lru.capacity {
		lru.removeLRU()
	}

	// Create new node and add to front
	newNode := &Node{key: key, val: val}
	lru.cache[key] = newNode
	lru.addToFront(newNode)
}

// moveToFront moves a node to the front (most recently used position)
func (lru *LruCache) moveToFront(node *Node) {
	lru.removeNode(node)
	lru.addToFront(node)
}

// addToFront inserts a node right after the dummy head
func (lru *LruCache) addToFront(node *Node) {
	node.next = lru.head.next
	node.prev = lru.head

	lru.head.next.prev = node
	lru.head.next = node
}

// removeNode removes a node from the linked list
func (lru *LruCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// removeLRU removes the least recently used item (right before tail)
func (lru *LruCache) removeLRU() {
	lruNode := lru.tail.prev
	if lruNode == lru.head {
		return // No nodes to remove
	}

	lru.removeNode(lruNode)
	delete(lru.cache, lruNode.key)
}
