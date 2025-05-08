package main

import "container/heap"

// https://leetcode.com/problems/process-tasks-using-servers/

func assignTasks(servers []int, tasks []int) []int {
	// [weight, index]
	freeServers := &PriorityQueue{}
	// [procTime, index]
	busyServers := &PriorityQueue{}
	heap.Init(freeServers)
	heap.Init(busyServers)

	for i, server := range servers {
		heap.Push(freeServers, [2]int{server, i})
	}

	res := make([]int, len(tasks))
	currTime := 0
	for i, task := range tasks {
		// ❗️Do not use currTime++
		// Test case: servers: [3,2,1]; tasks: [10,9,10,1,2,3]
		// At time t=10, multiple tasks (task 3, 4) are waiting
		// If you use curTime++, task 4 will be processed after task 3, which is incorrect.
		// They both should be processed at the same time when server is available.
		currTime = max(currTime, i)
		if freeServers.Len() == 0 {
			currTime = (*busyServers)[0][0]
		}

		// release busy server
		for busyServers.Len() > 0 && (*busyServers)[0][0] <= currTime {
			server := heap.Pop(busyServers).([2]int)
			id := server[1]
			weight := servers[id]
			heap.Push(freeServers, [2]int{weight, id})
		}

		server := heap.Pop(freeServers).([2]int)
		heap.Push(busyServers, [2]int{currTime + task, server[1]})
		res[i] = server[1]
	}

	return res
}

type PriorityQueue [][2]int

func (h PriorityQueue) Len() int { return len(h) }
func (h PriorityQueue) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
}
func (h PriorityQueue) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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
