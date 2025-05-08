package main

// https://leetcode.com/problems/top-k-frequent-elements/

import "math/rand"

// [1,1,1,2,2,3]
// [1,3] [2,3] [3,1]

// TC: O(n+m) average, worst case: O(n^2)
// SC: O(m)
// - m: the number of unique elements
func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	counts := make([][2]int, len(countMap))
	i := 0
	for num, count := range countMap {
		counts[i] = [2]int{count, num}
		i++
	}

	quickSelect(counts, 0, len(counts)-1, k)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = counts[i][1]
	}
	return res
}

func quickSelect(counts [][2]int, l, r, k int) {
	if l == r {
		return
	}

	n := r - l + 1
	pivot := counts[rand.Intn(n)+l][0]

	i, j := l, r

	for i < j {
		for counts[i][0] > pivot {
			i++
		}

		for counts[j][0] < pivot {
			j--
		}

		if i < j {
			counts[i], counts[j] = counts[j], counts[i]
			i++
			j--
		}
	}

	// l -> j
	// j+1 -> r
	if counts[j][0] < pivot {
		j--
	}

	if k <= j-l+1 {
		quickSelect(counts, l, j, k)
	} else {
		quickSelect(counts, j+1, r, k-(j-l+1))
	}
}
