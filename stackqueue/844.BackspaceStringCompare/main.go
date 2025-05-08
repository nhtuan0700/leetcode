package main

import "container/list"

// https://leetcode.com/problems/backspace-string-compare/

// same as description
// TC: O(n), SC: O(n)
func backspaceCompare1(s string, t string) bool {
	return process(s) == process(t)
}

func process(s string) string {
	stack := list.New()

	for _, c := range s {
		if c == '#' {
			if stack.Back() != nil {
				stack.Remove(stack.Back())
			}
			continue
		}
		stack.PushBack(c)
	}

	var res []rune
	for stack.Back() != nil {
		back := stack.Back()
		res = append(res, back.Value.(rune))
		stack.Remove(back)
	}
	return string(res)
}
