package main

// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	end := 0

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if end == len(strs[j]) {
				return string(strs[j])
			}
			if strs[j][i] != strs[0][i] {
				return string(strs[0][0:end])
			}
		}
		end++
	}

	return strs[0]
}
