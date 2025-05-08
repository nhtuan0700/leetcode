from typing import List

# https://leetcode.com/problems/prime-pairs-with-target-sum/description/

# 1. bruteforce
# i: 2 -> n - 2: candidate is if i is prime and n - i is prime
# TC: O(n*sqrt(n)*sqrt(n)) = O(n^2)

# 2. use sieve to generate prime number
# TC: O(nloglogn+m), SC: O(n + m)
# - m: number of prime number

# 3. use sieve to generate isPrime, then check from 2 -> n/2 + 1
# => optimal solution
# TC: O(sqrt(n)loglogn + n), SC: O(n)
class Solution:
	def findPrimePairs1(self, n: int) -> List[List[int]]:
		def isPrime(n):
			if n < 2:
				return False
			for i in range(2, int(n**0.5)+1):
				if n % i == 0:
					return False
			return True
		res = []
		for i in range(2, n//2+1):
			if isPrime(i) and isPrime(n - i):
				res.append([i, n-i])
		return res 
    
	def findPrimePairs2(self, n: int) -> List[List[int]]:
		if n < 4:
			return []
		isPrimes = [True] * n
		isPrimes[0] = isPrimes[1] = False
		primes = []
		for i in range(2, n):
			if isPrimes[i]:
				primes.append(i)
				for j in range(i*i, n, i):
					isPrimes[j] = False
		res = []
		for i in range(0, len(primes)):
			p = primes[i]
			if p > n // 2:
				break
			if isPrimes[n-p]:
				res.append([p, n - p])
		return res 

	def findPrimePairs3(self, n: int) -> List[List[int]]:
		if n < 4:
			return []
		isPrimes = [True] * n
		isPrimes[0] = isPrimes[1] = False
		for i in range(2, int(n**0.5) + 1):
			if isPrimes[i]:
				for j in range(i*i, n, i):
					isPrimes[j] = False
		res = []
		for i in range(2, n // 2 + 1):
			if isPrimes[i] and isPrimes[n-i]:
				res.append([i, n - i])
		return res 
