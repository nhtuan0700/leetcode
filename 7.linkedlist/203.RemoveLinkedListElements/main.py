# https://leetcode.com/problems/remove-linked-list-elements/description/

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

# TC: O(n), SC: O(1)
class Solution:
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        while head and head.val == val:
            head = head.next
        if not head:
            return None
        cur = head
        while cur:
            if cur.next and cur.next.val == val:
                cur.next = cur.next.next
            else:
                cur = cur.next
        return head
