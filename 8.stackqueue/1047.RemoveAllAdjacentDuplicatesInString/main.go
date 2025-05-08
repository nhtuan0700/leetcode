package main

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

import "container/list"

// TC: O(n), SC: O(n)
func removeDuplicates(s string) string {
	stack := list.New()

	for _, c := range s {
		if stack.Back() != nil {
			last := stack.Back().Value.(rune)
			if last == c {
				stack.Remove(stack.Back())
				continue
			}
			stack.PushBack(c)
			continue
		}
		stack.PushBack(c)
	}

	res := make([]rune, 0)
	for stack.Front() != nil {
		val := stack.Front().Value.(rune)
		res = append(res, val)
		stack.Remove(stack.Front())
	}

	return string(res)
}
