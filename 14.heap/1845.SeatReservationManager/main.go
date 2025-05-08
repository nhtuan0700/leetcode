package main

import "container/heap"

// https://leetcode.com/problems/non-overlapping-intervals/description/
// Sort the intervals by their end time. Then, iterate through each interval.
// - If the current interval overlaps with the previous one (i.e. start < last end), remove it.
// - This ensures we always keep the interval with the earliest end time, minimizing overlap.

// TC: O(nlogn)

type PriorityQueue []int

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i] < h[j] } // min heap
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PriorityQueue) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *PriorityQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeatManager struct {
	seats *PriorityQueue
}

func Constructor(n int) SeatManager {
	seats := &PriorityQueue{}
	heap.Init(seats)
	for i := 1; i <= n; i++ {
		heap.Push(seats, i)
	}

	return SeatManager{
		seats: seats,
	}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.seats).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.seats, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
