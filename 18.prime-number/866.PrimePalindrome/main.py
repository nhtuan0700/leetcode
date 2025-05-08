# https://leetcode.com/problems/prime-palindrome/description/

# generate palidrome number (p) 
# - Any even length palindrome must be divisble by 11, so we skip it
# - return if p >= n and isPrimeNumber(p)

# Idea: generate palindrome first, then check prime
# - because check whether it is palidrome in O(logn), and check whether it is prime in O(sqrt(n)). So we would probaly like to do the palindrome first

# O(p*sqrt(p)), worstcase O(10^5*sqrt(10^9))
class Solution:
    def isPrime(self, n: int) -> bool:
        if n < 2:
            return False
        for i in range(2, int(n**0.5)+1):
            if n % i == 0:
                return False
        return True

    def primePalindrome(self, n: int) -> int:
        if n >= 8 and n <= 11:
            return 11
        # 1 <= n <= 10^8
        # length of palidrome leftside is maximum 99999
        for i in range(2, int(1e5)):
            # reverse string and start from position 2
            p = int(str(i) + str(i)[-2::-1])
            if p >= n and self.isPrime(p):
                return p
        
        return -1
