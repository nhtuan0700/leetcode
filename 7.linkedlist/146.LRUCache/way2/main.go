package main

// https://leetcode.com/problems/lru-cache/description/
// TC: O(1), SC: O(3n)

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

type DLL struct {
	Head *Node
	Tail *Node
}

func NewDLL() DLL {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Previous = head
	return DLL{
		Head: head,
		Tail: tail,
	}
}

func (this *DLL) MoveToFront(node *Node) {
	// delete node
	node.Next.Previous = node.Previous
	node.Previous.Next = node.Next

	// move node to head
	head := this.Head
	node.Next = head.Next
	head.Next.Previous = node
	node.Previous = head
	head.Next = node
}

func (this *DLL) PushFront(value int) *Node {
	head := this.Head

	newNode := &Node{Value: value}
	newNode.Next = head.Next
	head.Next.Previous = newNode
	newNode.Previous = head
	head.Next = newNode

	return newNode
}

func (this *DLL) Remove(node *Node) {
	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous
}

type LRUCache struct {
	Lru      DLL // store key, head => recently used, tail => least used
	Capacity int
	Cache    map[int]int   // hashmap to store (key, value)
	Pos      map[int]*Node // hashmap to store (key, *Node)
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Lru:      NewDLL(),
		Capacity: capacity,
		Cache:    make(map[int]int),
		Pos:      make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.Cache[key]
	if !ok {
		return -1
	}

	this.UpdateLRU(key)
	return value
}

func (this *LRUCache) Put(key int, value int) {
	this.UpdateLRU(key)
	this.Cache[key] = value
}

func (this *LRUCache) UpdateLRU(key int) {
	node, ok := this.Pos[key]
	if !ok {
		if len(this.Cache) >= this.Capacity { // capacity is full, need to evict least recently used node
			// remove last node
			lastNode := this.Lru.Tail.Previous
			this.Lru.Remove(lastNode)
			delete(this.Cache, lastNode.Value)
			delete(this.Pos, lastNode.Value)
		}

		newNode := this.Lru.PushFront(key)
		this.Pos[key] = newNode
	} else {
		this.Lru.Remove(node)

		newNode := this.Lru.PushFront(key)
		this.Pos[key] = newNode
	}
}

/*
DLL: value, store the usage order
    - head: least recently used
    - tail: most recently used
map: key -> node of DLL
cur_capacity


get(key):
    - if key is not in map: return -1
    - else:
        - delete at map[key]
        - insert value to head
        - map[key] = head
        - return value

put(key, value)
    - if key is not in map:
        - if cur_capaicity < head
            - insert value to head
            - map[key] = head
        - else:
            - delete tail of DLL
            - insert value to head
            - map[key] = head
    - else:
        - delete at map[key]
        - insert value to head
        - map[key] = head


*/

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
