from typing import List

# https://leetcode.com/problems/maximum-prime-difference/

class Solution:
	# TC: O(n*sqrt(m))
    def maximumPrimeDifference(self, nums: List[int]) -> int:
        def isPrime(n) -> bool:
            if n < 2:
                return False
            for i in range(2, int(n**0.5)+1):
                if n % i == 0:
                    return False
            return True
        res = []
        for i, num in enumerate(nums):
            if isPrime(num):
                if not res:
                    res = [i, i]
                res[-1] = i
        return res[-1] - res[0]
    
	# TC: O(n), SC: O(1)
	# optimal solution
    def maximumPrimeDifference2(self, nums: List[int]) -> int:
        # generate isPrimes by using sieve of Eratosthenes
        def generateIsPrimes(n) -> List[bool]:
            n += 1
            isPrimes = [True] * n
            # 1, 0 is not prime number
            isPrimes[0] = isPrimes[1] = False
            for i in range(2, n):
                if isPrimes[i]:
                    for j in range(i*i, n, i):
                        isPrimes[j] = False
            return isPrimes
        isPrimes = generateIsPrimes(100)
        n = len(nums)
        for i in range(0, n):
            if isPrimes[nums[i]]:
                start = i
                break
        
        for i in range(n-1, -1, -1):
            if isPrimes[nums[i]]:
                return i - start
