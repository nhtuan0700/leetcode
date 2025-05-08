package main

// https://leetcode.com/problems/goal-parser-interpretation/

func interpret(command string) string {
	result := make([]byte, 0)

	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			result = append(result, command[i])
			continue
		}

		j := i
		for command[j] != ')' {
			j++
		}
		if string(command[i:j+1]) == "()" {
			result = append(result, 'o')
		} else if string(command[i:j+1]) == "(al)" {
			result = append(result, 'a', 'l')
		}

		i = j
	}

	return string(result)
}
