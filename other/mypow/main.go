package main

func myPow(x float64, n int) float64 {
	return pow(x, n, map[int]float64{})
}

// O(n)
func pow(x float64, n int, cached map[int]float64) float64 {
	if n == 0 {
		return 1
	}

	factor := x

	if n < 0 {
		factor = 1 / x
	}

	val, ok := cached[n]
	if ok {
		return val
	}

	if n%2 == 0 {
		cached[n] = pow(x, n/2, cached) * pow(x, n/2, cached)
	} else {
		cached[n] = factor * pow(x, n/2, cached) * pow(x, n/2, cached)
	}

	return cached[n]
}

// TC: O(logn), SC: O(logn)
func myPow2(x float64, n int) float64 {
	return pow2(x, n, map[int]float64{})
}

func pow2(x float64, n int, cached map[int]float64) float64 {
	if n == 0 {
		return 1
	}

	factor := x

	if n < 0 {
		factor = 1 / x
	}

	val, ok := cached[n]
	if ok {
		return val
	}

	half := pow(x, n/2, cached)
	if n%2 == 0 {
		cached[n] = half * half
	} else {
		cached[n] = factor * half * half
	}

	return cached[n]
}

// 11
// TC: O(logn), O(1)
func myPow3(x float64, n int) float64 {
	var result float64 = 1

	power := n
	base := x

	if power < 0 {
		power = -power
	}

	for power > 0 { // power=1
		if power%2 == 1 {
			result *= base // result= 1 * 2 * 4 * 16
		}
		base *= base // base = 2 * 2 * 4  * 16
		power /= 2 // power=0
	}

	if n < 0 {
		return 1 / result
	}

	return result
}
