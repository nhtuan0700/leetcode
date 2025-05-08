package main

// https://leetcode.com/problems/ransom-note/description/
// Map
// TC: O(n+m), SC: O(26) ~ O(1)
func canConstruct1(ransomNote string, magazine string) bool {
	freqMagazines := map[rune]int{}

	for _, c := range magazine {
		freqMagazines[c]++
	}

	for _, c := range ransomNote {
		if freqMagazines[c] == 0 {
			return false
		}

		freqMagazines[c]--
	}

	return true
}

// counting
// TC: O(n+m), SC: O(26) ~ O(1)
func canConstruct2(ransomNote string, magazine string) bool {
	// ransomNote and magazine consist of lowercase English letters.
	counts := make([]int, 26)

	for _, c := range magazine {
		counts[c-'a']++
	}

	for _, c := range ransomNote {
		if counts[c-'a'] == 0 {
			return false
		}

		counts[c-'a']--
	}

	return true
}
