package main

func isAlphanumeric(c byte) bool {
	if '0' <= c && c <= '9' {
		return true
	}
	if 'A' <= c && c <= 'Z' {
		return true
	}
	if 'a' <= c && c <= 'z' {
		return true
	}
	return false
}

func toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c - 'A' + 'a'
	}

	return c
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		for left < len(s)-1 && !isAlphanumeric(s[left]) {
			left++
		}
		if left == len(s)-1 {
			break
		}

		for right > 0 && !isAlphanumeric(s[right]) {
			right--
		}
		if right == 0 {
			break
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--

	}

	return true

}
