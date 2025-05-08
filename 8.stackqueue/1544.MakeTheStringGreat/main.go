package main

// https://leetcode.com/problems/make-the-string-great/

// TC: O(n), SC: O(n)
func makeGood(s string) string {
	stack := []rune{}

	for _, c := range s {
		if len(stack) > 0 && isBad(c, stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return string(stack)
}

func isBad(c1, c2 rune) bool {
	return abs(int(c1 - c2)) == 32
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
