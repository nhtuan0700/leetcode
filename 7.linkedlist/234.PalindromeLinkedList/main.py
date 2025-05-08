from typing import Optional

# https://leetcode.com/problems/palindrome-linked-list/

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
# TC: O(n), SC: O(1)
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        slow = head
        fast = head
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next
        cur = slow
        # reverse
        head2 = None
        while cur:
            _next = cur.next
            cur.next = head2
            head2 = cur
            cur = _next
        
        while head and head2:
            if head.val != head2.val:
                return False
            head = head.next
            head2 = head2.next

        return True
