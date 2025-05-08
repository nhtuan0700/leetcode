from typing import List

# https://leetcode.com/problems/closest-prime-numbers-in-range/

# Use sieve of Eratosthenes to generate prime numbers from left to right
# The find min difference value between 2 element consecutive
# TC: O(nloglogn) SC: O(n)
# - n: right
class Solution:
    def closestPrimes(self, left: int, right: int) -> List[int]:
        n = right + 1
        isPrimes = [True] * n
        isPrimes[0:2] = [False] * 2
        for i in range(2, int(n**0.5) + 1):
            if isPrimes[i]:
                for j in range(i*i, n, i):
                    isPrimes[j] = False
        primes = [i for i in range(left, n) if isPrimes[i]]
        if len(primes) < 2:
            return [-1, -1]
        res = []
        for i in range(0, len(primes) - 1):
            if not res or primes[i+1] - primes[i] < res[1] - res[0]:
                res = [primes[i], primes[i+1]]
        return res
    
# if right is large (10^9), but the range [left, right] is small, we could optimize this by:
# 1. Generate prime numbers from 2 -> sqrt(right)
# 2. In range (left, right), mark not prime number from prime number list in step 1
# 3. From step 2, we could get the prime number from left to right
# TC: O(nloglogn + mloglogn), SC: O(n+m)
# - n: sqrt(right), m: right-left
class Solution2:
    def closestPrimes(self, left: int, right: int) -> List[int]:
        n = int(right**0.5)+1
        isPrimes = [True] * n
        primes = []
        for i in range(2, n):
            if isPrimes[i]:
                primes.append(i)
                for j in range(i*i, n, i):
                    isPrimes[j] = False
        
        n2 = right - left + 1
        isPrimes2 = [True] * n2
        if left == 1: isPrimes2[0] = False
        primes2 = []
        for p in primes:
            # find first product of the prime number
            first = ((left + p - 1) // p) * p
            if first == p:
                first += p
            for j in range(first, right + 1, p):
                isPrimes2[j - left] = False

        for i in range(left, right+1):
            if isPrimes2[i - left]:
                primes2.append(i)
        num1 = num2 = -1
        for i in range(0, len(primes2) - 1):
            if num1 == -1 or (primes2[i+1] - primes2[i] < num2 - num1):
                num1 = primes2[i]
                num2 = primes2[i+1]
        return [num1, num2]
