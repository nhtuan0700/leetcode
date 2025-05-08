package main

import "container/heap"

// TC: O(nlogk), SC: O(k)
func kClosest(points [][]int, k int) [][]int {
	maxHeap := &PriorityQueue{}
	for i, p := range points {
		distance := p[0]*p[0] + p[1]*p[1]
		heap.Push(maxHeap, [2]int{distance, i})
		if maxHeap.Len() > k {
			heap.Pop(maxHeap)
		}
	}

	res := make([][]int, k)
	i := 0
	for maxHeap.Len() > 0 {
		val := heap.Pop(maxHeap).([2]int)
		res[i] = points[val[1]]
		i++
	}

	return res
}

type PriorityQueue [][2]int

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}

func (h *PriorityQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
