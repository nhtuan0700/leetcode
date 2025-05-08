package main

// we will calulate until total == 1 or meet the pre-calculated total

func runIsHappy(n int, cached map[int]bool) bool {
	total := 0

	for n > 0 {
		total += (n % 10) * (n % 10)
		n /= 10
	}

	if total == 1 {
		return true
	}

	if cached[total] {
		return false
	}

	cached[total] = true
	return runIsHappy(total, cached)
}

func isHappy(n int) bool {
	return runIsHappy(n, make(map[int]bool))
}
