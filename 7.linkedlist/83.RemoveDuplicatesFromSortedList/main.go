package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// TC: O(n)
// 1 -> 1 -> 2
func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next // delete next node
		}
		current = current.Next
	}

	return head
}
