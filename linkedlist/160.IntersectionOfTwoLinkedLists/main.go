package main

// https://leetcode.com/problems/intersection-of-two-linked-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

// TC: O(n+m), SC: O(n)
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	cached := map[*ListNode]bool{}
	for headA != nil {
		cached[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if cached[headB] {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

// convert to negative number
// TC: O(n+m), SC: O(1)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	tempA := headA
	for headA != nil {
		headA.Val = -headA.Val
		headA = headA.Next
	}

	for headB != nil {
		if headB.Val < 0 {
			break
		}
		headB = headB.Next
	}

	for tempA != nil {
		tempA.Val = -tempA.Val
		tempA = tempA.Next
	}

	return headB
}

// use 2 pointers
// n: length of headA
// m: length of headB
// n > m
// - true: headA start from (n-m)
// - false: headB start from (m-n)
// Then jump both headA & headB until headA = = headB
// TC: O(n+m), SC: O(1)
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	n := findLength(headA)
	m := findLength(headB)
	for n > m {
		headA = headA.Next
		n--
	}
	for m > n {
		headB = headB.Next
		m--
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

func findLength(head *ListNode) int {
	n := 0
	current := head
	for current != nil {
		n++
		current = current.Next
	}

	return n
}
