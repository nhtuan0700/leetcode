# https://leetcode.com/problems/happy-number/description/

# TC: O(logn), SC: O(1)
class Solution:
    def isHappy(self, n: int) -> bool:
        visits = set()

        # sumSquare runs in O(log n) because there are up to log10(n) digits
        def sumSquare(n):
            res = 0
            while n > 0:
                res += (n % 10) ** 2
                n //= 10
            return res

        # total number of unique values from sumSquare(n) is at most 243
        def helper(n):
            nonlocal visits
            if n == 1:
                return True
            if n in visits:
                return False
            visits.add(n)
            return helper(sumSquare(n))

        return helper(n)
