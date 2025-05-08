package main

// https://leetcode.com/problems/shuffle-string/

// [c,o,d,e,l,e,e,t]
// [4,5,6,7,0,2,1,3]

// [_,_,_,_,_,_,_,_]

func restoreString(s string, indices []int) string {
	result := make([]byte, len(s))
	for i := 0; i < len(result); i++ {
		result[indices[i]] = s[i]
	}

	return string(result)
}
