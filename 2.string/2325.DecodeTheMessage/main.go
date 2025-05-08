package main

// https://leetcode.com/problems/decode-the-message/

// Use the first appearance of all 26 lowercase English letters in key as the order of the substitution table.
// map[key[i]] = alpha[i]

func decodeMessage(key string, message string) string {
	charMap := map[byte]byte{}
	result := make([]byte, 0)

	c := byte('a')
	for i := 0; i < len(key); i++ {
		if c > 'z' {
			break
		}
		if key[i] == ' ' {
			continue
		}
		if _, ok := charMap[key[i]]; ok {
			continue
		}
		charMap[key[i]] = c
		c += byte(1)
	}

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c != ' ' {
			c = charMap[c]
		}

		result = append(result, c)
	}
	return string(result)
}
