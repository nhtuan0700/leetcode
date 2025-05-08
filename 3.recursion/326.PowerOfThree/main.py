# https://leetcode.com/problems/power-of-three/

# Nếu a là một số nguyên dương và n = a^k là số nguyên dương lớn nhất mà vẫn là lũy thừa của a và ≤ MAX_INT, 
# thì mọi lũy thừa của a nhỏ hơn hoặc bằng n đều là ước của n.
# TC: O(1), SC: O(1)
class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        # return n > 0 and 1_162_261_467 % n == 0
        return n > 0 and 3**19 % n == 0

# TC: O(logn), SC: O(logn)
class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        if n <= 0:
            return False
        if n == 1:
            return True

        return n%3 == 0 and self.isPowerOfThree(n//3)

# TC: O(logn), SC: O(1)
class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        if n <= 0:
            return False
        while n % 3 == 0:
            n /= 3
        
        return n == 1
