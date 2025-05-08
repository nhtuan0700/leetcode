package main

// https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/description/?envType=daily-question&envId=2025-06-03

import "container/list"

// TC: O(V + E), SC: O(E)
func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	n := len(containedBoxes)
	vis := make([]bool, n)
	opens := make([]bool, n)
	res := 0

	q := list.New()
	for _, b := range initialBoxes {
		q.PushBack(b)
		vis[b] = true
	}

	for q.Front() != nil {
		v := q.Front().Value.(int)
		q.Remove(q.Front())

		if !opens[v] && status[v] == 1 {
			res += candies[v]
			opens[v] = true

			for _, k := range keys[v] {
				status[k] = 1
				if vis[k] && !opens[k] {
					q.PushBack(k)
				}
			}
		}

		for _, u := range containedBoxes[v] {
			if !vis[u] {
				q.PushBack(u)
				vis[u] = true
			}
		}
	}

	return res
}
