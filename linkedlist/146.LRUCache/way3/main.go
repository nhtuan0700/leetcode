package main

import (
	"container/list"
)

// https://leetcode.com/problems/lru-cache/description/
// TC: O(1), SC: O(3n)

type LRUCache struct {
	Lru      *list.List // store int[key, value], head => recently used, tail => least used
	Capacity int
	Cache    map[int]*list.Element // hashmap to store key and value
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Lru:      list.New(),
		Capacity: capacity,
		Cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}

	this.Lru.MoveToFront(node)
	return node.Value.([]int)[1]
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Cache[key]
	if !ok {
		if len(this.Cache) == this.Capacity { // capacity is full, need to evict least recently used node
			// remove last node
			lastNode := this.Lru.Back().Prev()
			this.Lru.Remove(lastNode)
			delete(this.Cache, lastNode.Value.([2]int)[0])
		}
		newNode := this.Lru.PushFront([2]int{key, value})
		this.Cache[key] = newNode
	} else {
		node.Value = [2]int{key, value}
		this.Lru.MoveToFront(node)
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
