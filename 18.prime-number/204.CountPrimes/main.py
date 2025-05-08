# https://leetcode.com/problems/count-primes/

# Idea:
# - Generate primes number and count -> use sieve of Eratosthenes

# TC: O(nloglogn), SC: O(n)
class Solution:
    def countPrimes(self, n: int) -> int:
        if n <= 2:
            return 0
        isPrimes = [1] * n
        isPrimes[0:2] = [0] * 2
        for i in range(2, int(n**0.5)+1):
            if isPrimes[i]:
                for j in range(i*i, n, i):
                    isPrimes[j] = 0
        return sum(isPrimes)


# Check only odd number (2 is counted separately)
# use sieve of Eratosthenes to mark odd number is not prime number
# TC: O(nloglogn), SC: O(n)
class Solution2:
    def countPrimes(self, n: int) -> int:
        if n <= 2:
            return 0
        isPrimes = [1] * n
        cnt = 1 # include 2
        for i in range(3, n, 2):
            if isPrimes[i]:
                cnt += 1
                for j in range(i*i, n, 2*i):
                    isPrimes[j] = 0
        return cnt
