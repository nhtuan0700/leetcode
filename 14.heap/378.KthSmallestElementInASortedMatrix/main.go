package main

// Pattern
// Mục tiêu	  | Dùng Heap gì	        | Vì sao?
// K Smallest |	Max Heap (k phần tử). | Vì ta giữ k phần tử nhỏ nhất ⇒ bỏ phần tử lớn nhất khi dư
// K Largest	| Min Heap (k phần tử)	| Vì ta giữ k phần tử lớn nhất ⇒ bỏ phần tử nhỏ nhất khi dư

// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

import "container/heap"

// TC: O(n^2logk) SC: O(k)
// TODO?
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	maxHeap := &PriorityQueue{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			heap.Push(maxHeap, matrix[i][j])

			if maxHeap.Len() > k {
				heap.Pop(maxHeap)
			}
		}
	}

	return heap.Pop(maxHeap).(int)
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
