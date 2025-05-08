package main

// https://leetcode.com/problems/smallest-number-in-infinite-set/

import "container/heap"

// TC: O(log n), SC: O(n)
// n: size of heap
type SmallestInfiniteSet struct {
	smallestNum int
	pq          *PriorityQueue
	set         map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		smallestNum: 1,
		pq:          &PriorityQueue{},
		set:         make(map[int]bool),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.pq.Len() > 0 {
		val := heap.Pop(this.pq).(int)
		delete(this.set, val)
		return val
	}

	num := this.smallestNum
	this.smallestNum++
	return num
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.smallestNum && !this.set[num] {
		heap.Push(this.pq, num)
		this.set[num] = true
	}
}

type PriorityQueue []int

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i] < h[j] }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *PriorityQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
