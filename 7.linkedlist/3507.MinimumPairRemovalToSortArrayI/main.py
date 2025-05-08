from collections import List

# https://leetcode.com/problems/minimum-pair-removal-to-sort-array-i/

class Node:
    def __init__(self, value, next_ = None):
        self.value = value
        self.next = next_

# TC: O(n^2), SC: O(n)
class Solution:
    def minimumPairRemoval(self, nums: List[int]) -> int:
        head = Node(nums[0])
        cur = head
        for i in range(1, len(nums)):
            cur.next = Node(nums[i])
            cur = cur.next
        
        count = 0
        while head.next:
            minTotal = head.value + head.next.value
            pairNode = head
            cur = head
            desc = False
            while cur.next:
                if cur.value > cur.next.value:
                    desc = True
                if minTotal > cur.value + cur.next.value:
                    pairNode = cur
                    minTotal = cur.value + cur.next.value
                cur = cur.next
            if not desc:
                break
            pairNode.value = pairNode.value + pairNode.next.value
            pairNode.next = pairNode.next.next
            count += 1
        return count
