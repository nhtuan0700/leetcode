package main

import "sort"

// https://leetcode.com/problems/group-anagrams/submissions/1637933044/

// n: strs length
// k: strs[i].length
// TC: O(nklogk). Worst case: 100*10^4*log(10^4) = 14,000,000
// loop str: strs:
//  - add maps: map[sorted(str)] = []string{...}

func groupAnagrams(strs []string) [][]string {
	maps := make(map[string][]string)

	for _, str := range strs {
		s := []byte(str)
		sort.SliceStable(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		maps[string(s)] = append(maps[string(s)], str)
	}

	result := make([][]string, len(maps))
	i := 0
	for _, strs := range maps {
		result[i] = strs
		i++
	}

	return result
}
