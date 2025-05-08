package main

import "container/heap"

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/

// O(nlogk)

func findKthLargest(nums []int, k int) int {
	pq := &PriorityQueue{}
	for _, num := range nums {
		heap.Push(pq, num)
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	return heap.Pop(pq).(int)
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
