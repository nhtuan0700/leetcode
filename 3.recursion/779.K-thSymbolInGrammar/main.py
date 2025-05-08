# 1: 0
# 2: 01
# 3: 0110
# 4: 01101001
# ...
# n: 2^(n-1)

class Solution:
    def kthGrammar(self, n: int, k: int) -> int:
        if n == 1:
            return 0
        l = 1 << (n-1)
        m = l // 2
        if k <= m:
            return self.kthGrammar(n-1, k)
        return 1-self.kthGrammar(n-1, k-m)

class Solution2:
    def kthGrammar(self, n: int, k: int) -> int:
        if n == 1:
            return 0
        
        res = 0
        while n > 1:
            l = 1 << (n-1)
            m = l // 2
            if k > m:
                k -= m
                res = 1 - res
            n -= 1
        
        return res
