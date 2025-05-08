package main

import "sort"

// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/

// max TC: O(nlogn)
// [5,1,3], [1,2,3,4,5]
// sort potions: [1,2,3,4,5] // O(mlogm)
// loop n (On) => n(logm)
// bsearch to find success // O(logm)
// => TC: O((m + n)logm), SP: O(1)

// similar:  largest num (idex) < target, m - index => success
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(potions)

	for i, sp := range spells {
		// O(logm)
		index := getCountOfPairProductSmallerThanTarget(sp, potions, success)
		spells[i] = m - (index + 1)
	}

	return spells
}

func getCountOfPairProductSmallerThanTarget(spell int, potions []int, success int64) int {
	l := 0
	h := len(potions) - 1

	for l <= h {
		m := l + (h-l)/2
		if int64(potions[m]*spell) < success {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return h
}
