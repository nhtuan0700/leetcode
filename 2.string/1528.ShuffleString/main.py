from typing import List

# TC: O(n), SC: O(n)
class Solution:
    def restoreString(self, s: str, indices: List[int]) -> str:
        n = len(s)
        res = [''] * n
        for i in range(n):
            res[indices[i]] = s[i]

        return ''.join(res)
