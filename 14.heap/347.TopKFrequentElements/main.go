package main

import "container/heap"

// https://leetcode.com/problems/top-k-frequent-elements/

// TC: O(n + mlogk), SC: O(m+k) ~ SC: O(m)
// m: number of unique nums
func topKFrequent(nums []int, k int) []int {
	// map[num]count
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	// [count, num]
	pq := &PriorityQueue{}
	for num, count := range counts {
		heap.Push(pq, [2]int{count, num})

		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	res := make([]int, k)
	i := 0
	for i < k {
		val := heap.Pop(pq).([2]int)
		res[i] = val[1]
		i++
	}
	return res
}

type PriorityQueue [][2]int

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i][0] < h[j][0] }
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
