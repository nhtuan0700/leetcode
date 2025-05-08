# https://leetcode.com/problems/intersection-of-two-linked-lists

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# TC: O(n+m), SC: O(1)
# m: length of list A
# n: length of list B
class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        # reverse
        lenA = self.findLength(headA)
        lenB = self.findLength(headB)
        while lenA > lenB:
            headA = headA.next
            lenA -= 1
        while lenB > lenA:
            headB = headB.next
            lenB -= 1
        while headA != headB:
            headA = headA.next
            headB = headB.next
        return headA
    def findLength(self, head):
        n = 0
        cur = head
        while cur:
            n += 1
            cur = cur.next
        return n
