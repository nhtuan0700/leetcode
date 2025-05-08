# https://leetcode.com/problems/valid-anagram/

# TC: O(n), SC: O(max(0,26)) = O(1)
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        freqs = {}
        for c in s:
            if c not in freqs:
                freqs[c] = 1
            else:
                freqs[c] += 1
        
        for c in t:
            if not freqs.get(c):
                return False
            if freqs[c] == 1:
                del freqs[c]
                continue
            freqs[c] -= 1
        
        return True
