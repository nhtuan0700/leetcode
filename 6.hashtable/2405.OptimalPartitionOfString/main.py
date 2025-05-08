# https://leetcode.com/problems/optimal-partition-of-string/

# TC: O(n), SC: O(1)
class Solution:
    def partitionString(self, s: str) -> int:
        freqs = set()
        cnt = 0
        for c in s:
            if c in freqs:
                cnt += 1
                freqs = set()
            freqs.add(c)
        return cnt + 1
        