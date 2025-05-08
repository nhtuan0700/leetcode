from collections import defaultdict

# TC: O(n), SC: O(k)
# n = highLimit - lowLimit + 1
# k <= 9 * digit_of_highLimit
class Solution:
    def countBalls(self, lowLimit: int, highLimit: int) -> int:
        def sumDigit(n):
            total = 0
            while n != 0:
                total += n%10
                n //= 10
            return total
        boxes = defaultdict(int)
        res = 0
        for i in range(lowLimit, highLimit + 1):
            key = sumDigit(i)
            boxes[key] += 1
            res = max(res, boxes[key])
            
        return res
