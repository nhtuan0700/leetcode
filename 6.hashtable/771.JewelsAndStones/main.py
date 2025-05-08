# https://leetcode.com/problems/jewels-and-stones/

# TC: O(n+m), SC: O(m)
# n: length stone
# m: len jewels
class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        jewelDict = set()
        for jw in jewels:
            jewelDict.add(jw)
        res = 0
        for st in stones:
            if st in jewelDict:
                res += 1
        return res
