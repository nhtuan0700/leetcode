# https://leetcode.com/problems/fibonacci-number/

# TC: O(n), SC: O(n)
class Solution:
    def fib(self, n: int) -> int:
        if n <= 1:
            return n
        fibs = [0,1]
        for i in range(2, n+1):
            fibs.append(fibs[i-1] + fibs[i-2])
        return fibs[n]
    
# dynmaic programing
# TC: O(n): SC: O(n)
class Solution:
    def fib(self, n: int) -> int:
        dp = [-1] * (n + 1)
        def _fib(n):
            nonlocal dp
            if n <= 1:
                return n
            if dp[n] != -1:
                return dp[n]
            dp[n] = _fib(n-1) + _fib(n-2)
            print(dp)
            return dp[n]
        return _fib(n)
    