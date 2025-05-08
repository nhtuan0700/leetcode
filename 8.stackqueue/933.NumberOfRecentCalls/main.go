package main

// https://leetcode.com/problems/number-of-recent-calls/

import "container/list"

type RecentCounter struct {
	queue *list.List
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: list.New(),
	}
}

// TC: O(1) amortized, SC: O(max(3000))
// - worst case: 3000 remove operations but not always
// - n pings -> max 2n operations => for each ping time, we will have the average O(1)
// return counter elements in range [t - 3000, t] <=> first >= t - 3000
func (this *RecentCounter) Ping(t int) int {
	this.queue.PushBack(t)
	for true {
		first := this.queue.Front().Value.(int) // O(1)
		// if first + 3000 >= t {
		// if t - first <= 3000 {
		if first >= t-3000 {
			break
		}
		this.queue.Remove(this.queue.Front()) // O(1)
	}

	return this.queue.Len()
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
