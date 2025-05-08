package main

// RLRRLLRLRL

func balancedStringSplit(s string) int {
	r := 0
	l := 0
	count := 1

	for _, c := range s {
		if r > 0 && l > 0 && l == r {
			count++
			r = 0
			l = 0
		}

		if c == 'R' {
			r++
		}

		if c == 'L' {
			l++
		}
	}

	return count
}
