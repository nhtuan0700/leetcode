package main

// https://leetcode.com/problems/optimal-partition-of-string/

// hash table
// s = "abacaba"
// loop c: s
//  map[c] -> result++,
//  else clear map

// TC: O(n), SC: O(n)
func partitionString(s string) int {
	freqs := map[rune]bool{}

	result := 0
	for _, r := range s {
		if freqs[r] {
			freqs = map[rune]bool{}
			result++
		}
		freqs[r] = true

	}

	return result + 1
}

// using bit operator
// TC: O(n), SC: O(1)
func partitionString2(s string) int {
	mask, count := 0, 0

	for _, c := range s {
		bit := 1 << (c - 'a') // convert c into bit
		if mask&bit != 0 {    // check if c is existed
			count++
			mask = bit // start new c
		} else {
			mask |= bit // store c
		}
	}

	return count + 1
}
