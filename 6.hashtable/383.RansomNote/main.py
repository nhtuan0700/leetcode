# https://leetcode.com/problems/ransom-note/

class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        letters = {}
        for c in magazine:
            letters[c] = letters.get(c, 0) + 1
        
        for c in ransomNote:
            if letters.get(c, 0) <= 0:
                return False
            letters[c] -= 1
        
        return True
