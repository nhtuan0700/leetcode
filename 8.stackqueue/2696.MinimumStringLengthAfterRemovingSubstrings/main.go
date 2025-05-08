package main

// https://leetcode.com/problems/minimum-string-length-after-removing-substrings/
func minLength(s string) int {
	stack := []rune{}

	for _, c := range s {
		if (c == 'B' && len(stack) > 0 && stack[len(stack)-1] == 'A') ||
			(c == 'D' && len(stack) > 0 && stack[len(stack)-1] == 'C') {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, c)
	}

	return len(stack)
}
