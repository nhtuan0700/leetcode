package main

// https://leetcode.com/problems/super-pow/

// idea
// 7^12345 = 7^10000 * 7^2000 * 7^300 * 7^40 * 7^5
// = (((((7^10)^10)^10)^10)^1 * ((7^10)^10)^10^4) * (((7^10)^10)^3) * (7^10)^4 * 7^5

// properties:
// (a*b) mod n = [(a mod n) * (b mod n)] mod n
// TC: O(nlogd) ~ O(n), SC: O(logd) ~ O(1)
// d: digit (max: 9)
func superPow(a int, b []int) int {
	result := 1
	mod := 1337

	for i := len(b) - 1; i >= 0; i-- { // O(n)
		d := b[i]
		result = (result * myPow(a, d, mod)) % mod // O(logd)
		a = myPow(a, 10, mod)                      // O(log10)
	}

	return result
}

func myPow(base, exp, mod int) int {
	if exp == 0 {
		return 1
	}

	half := myPow(base, exp/2, mod)
	result := (half * half) % mod

	if exp%2 == 1 {
		result = (result * base) % mod
	}

	return result
}
