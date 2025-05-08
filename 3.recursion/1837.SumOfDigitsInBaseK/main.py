class Solution:
    def sumBase(self, n: int, k: int) -> int:
        if n == 0:
            return 0
        return self.sumBase(n//k, k) + n%k
