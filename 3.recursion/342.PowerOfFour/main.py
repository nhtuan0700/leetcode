# TC: O(log4(n)) ~ O(logn) SC: O(logn)
class Solution:
    def isPowerOfFour(self, n: int) -> bool:
        if n < 1:
            return False
        if n == 1:
            return True
        return n % 4 == 0 and self.isPowerOfFour(n//4)
    

class Solution:
    def isPowerOfFour(self, n: int) -> bool:
        if n < 1:
            return False
        while n % 4 == 0:
            n //= 4
        
        return n == 1
