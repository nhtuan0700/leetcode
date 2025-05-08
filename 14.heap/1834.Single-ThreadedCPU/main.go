package main

// https://leetcode.com/problems/single-threaded-cpu/

import (
	"container/heap"
	"sort"
)

// TC: O(nlogn) SC: O(n)
func getOrder(tasks [][]int) []int {
	n := len(tasks)
	arr := make([][3]int, n)
	for i, t := range tasks {
		arr[i] = [3]int{t[0], t[1], i}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	res := make([]int, 0)
	pq := &PriorityQueue{}
	i := 0
	t := 0
	heap.Init(pq)

	for len(res) < n {
		for i < n && arr[i][0] <= t {
			procTime, id := arr[i][1], arr[i][2]
			heap.Push(pq, Task{procTime, id}) // O(log k)
			i++
		}

		if pq.Len() > 0 {
			task := heap.Pop(pq).(Task) // O(log k)
			procTime, id := task.ProcTime, task.Index
			res = append(res, id)
			t += procTime
		} else {
			t = arr[i][0]
		}
	}

	return res
}

type Task struct {
	ProcTime int
	Index    int
}

type PriorityQueue []Task

func (h PriorityQueue) Len() int { return len(h) }

// minheap
// procTime -> index
func (h PriorityQueue) Less(i, j int) bool {
	if h[i].ProcTime == h[j].ProcTime {
		return h[i].Index < h[j].Index
	}
	return h[i].ProcTime < h[j].ProcTime
}
func (h PriorityQueue) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x any) {
	*h = append(*h, x.(Task))
}

func (h *PriorityQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
