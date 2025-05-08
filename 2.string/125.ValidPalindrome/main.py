# https://leetcode.com/problems/valid-palindrome/

# TC: O(n), SC: O(1)
class Solution:
    def isPalindrome(self, s: str) -> bool:
        n = len(s)
        l = 0
        r = n - 1
        while l < r:
            while l < n -1 and not s[l].isalnum():
                l += 1
            if l == n - 1:
                break
            
            while r > 0 and not s[r].isalnum():
                r -= 1
            if r == 0:
                break
            if s[l].lower() != s[r].lower():
                return False
            l += 1
            r -= 1
        return True
