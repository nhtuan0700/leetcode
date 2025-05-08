from typing import List

# https://leetcode.com/problems/distinct-prime-factors-of-product-of-array/
 
# TC: O(n*sqrt(max(nums)))
class Solution:
    def distinctPrimeFactors(self, nums: List[int]) -> int:
        primes = set()
        def factorize(n):
            nonlocal primes
            i = 2
            while i * i <= n:
                while n % i == 0:
                    primes.add(i)
                    n = n // i
                i += 1
            if  n > 1:
                primes.add(n)

        for num in nums:
            factorize(num, primes)

        return len(primes)
