package main

// TC: O(n), SC: O(n)
func numJewelsInStones1(jewels string, stones string) int {
	jewelMaps := map[byte]bool{}
	for _, je := range jewels {
		jewelMaps[byte(je)] = true
	}

	result := 0
	for _, st := range stones {
		if jewelMaps[byte(st)] {
			result++
		}
	}

	return result
}

// TC: O(n), SC: O(n)
func numJewelsInStones2(jewels string, stones string) int {
	jewelMaps := map[byte]int{}
	for _, je := range jewels {
		jewelMaps[byte(je)] = 0
	}

	for _, st := range stones {
		if _, ok := jewelMaps[byte(st)]; ok {
			jewelMaps[byte(st)]++
		}
	}

	result := 0
	for _, count := range jewelMaps {
		result += count
	}

	return result
}
