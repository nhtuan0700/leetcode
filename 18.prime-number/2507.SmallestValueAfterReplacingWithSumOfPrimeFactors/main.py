# https://leetcode.com/problems/smallest-value-after-replacing-with-sum-of-prime-factors/

# TC: O(sqrt(n)*logn), SC: O(1)
class Solution:
	def smallestValue1(self, n: int) -> int:
		def factorizeAndSum(n):
			i = 2
			factors = []
			while i * i <= n:
				while n % i == 0:
					factors.append(i)
					n = n // i
				i += 1
			if n > 1: factors.append(n)
			return sum(factors)
		while True:
			v = factorizeAndSum(n)
			if n == v:
				break
			n = v
		return n

	def smallestValue2(self, n: int) -> int:
		def factorizeAndSum(n):
			factors = []
			while n % 2 == 0:
				factors.append(2)
				n //= 2
			i = 3
			while i * i <= n:
				while n % i == 0:
					factors.append(i)
					n //= i
				i += 2
			if n > 1: factors.append(n)
			return sum(factors)
		while True:
			v = factorizeAndSum(n)
			if n == v:
				break
			n = v
		return n 

	def smallestValue3(self, n: int) -> int:
		def factorizeAndSum(n):
			total = 0
			while n % 2 == 0:
				total += 2
				n //= 2
			i = 3
			while i * i <= n:
				while n % i == 0:
					total += i
					n //= i
				i += 2
			if n > 1: total += n
			return total
		while True:
			v = factorizeAndSum(n)
			if n == v:
				break
			n = v
		return n 
        