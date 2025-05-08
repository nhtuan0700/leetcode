package main

// https://leetcode.com/problems/power-of-three/

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}

	if n == 0 || n%3 != 0 {
		return false
	}

	return isPowerOfThree(n / 3)
}

func isPowerOfThree2(n int) bool {
	for n > 0 && n != 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}

	return n == 1
}
