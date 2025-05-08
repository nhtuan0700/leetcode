package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runeMap := map[rune]int{}
	for _, c := range s {
		runeMap[c]++
	}

	for _, c := range t {
		if runeMap[c] == 0 {
			return false
		}

		runeMap[c]--
	}

	return true
}
