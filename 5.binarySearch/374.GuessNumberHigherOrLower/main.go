package main

// https://leetcode.com/problems/guess-number-higher-or-lower/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int { return 0 }

func guessNumber1(n int) int {
	left := 1
	right := n

	for left <= right {
		mid := left + (right-left)/2
		// pick < guess
		if guess(mid) == -1 {
			right = mid - 1
		} else if guess(mid) == 1 { // pick > guess
			left = mid + 1
		} else { // pick = guess
			return mid
		}
	}

	return -1
}

func guessNumber2(n int) int {
	left := 1
	right := n

	for left <= right {
		mid := left + (right-left)/2
		// pick = guess
		if guess(mid) == 0 {
			return mid
		}
		// pick < guess
		if guess(mid) == -1 {
			right = mid - 1
		} else { // pick > guess
			left = mid + 1
		}
	}

	return -1
}
