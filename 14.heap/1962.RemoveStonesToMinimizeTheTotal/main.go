package main

// https://leetcode.com/problems/remove-stones-to-minimize-the-total/description/

import "container/heap"

// TC: O(nlogn + klogn), SC: O(n)
func minStoneSum(piles []int, k int) int {
	maxHeap := &PriorityQueue{}
	total := 0
	for _, p := range piles {
		total += p
		heap.Push(maxHeap, p)
	}

	for k > 0 {
		cur := heap.Pop(maxHeap).(int)
		// ceil
		removed := (cur + 2 - 1) / 2
		total = total - (cur - removed)
		heap.Push(maxHeap, removed)
		k--
	}

	return total
}

type PriorityQueue []int

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i] > h[j] }
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
