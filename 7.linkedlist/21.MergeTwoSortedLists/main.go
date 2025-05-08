package main

// https://leetcode.com/problems/merge-two-sorted-lists/

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

// n: size of list1
// m: size of list2
// TC: O(n + m)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	currentNode := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			currentNode.Next = list1
			list1 = list1.Next
		} else {
			currentNode.Next = list2
			list2 = list2.Next
		}
		currentNode = currentNode.Next
	}

	if list1 != nil {
		currentNode.Next = list1
	} else {
		currentNode.Next = list2
	}

	return head.Next
}
