# TC: O(logn), SC: O(logn)
class Solution:
    def convertToBase7(self, num: int) -> str:
        if int(num/7) == 0:
            return str(num)
        mod = num % 7
        if mod < 0:
            mod *= -1
        if num < 0 and mod > 0:
            mod = 7 - mod
        return self.convertToBase7(int(num/7)) + str(mod)

# TC: O(logn), SC: O(logn)
class Solution:
    def convertToBase7(self, num: int) -> str:
        sign = ""
        if num < 0:
            num *= -1
            sign = "-"
        if num // 7 == 0:
            return sign + str(num)
        return sign + self.convertToBase7(num//7) + str(num%7)

# TC: O(logn), SC: O(1)
class Solution:
    def convertToBase7(self, num: int) -> str:
        if num == 0:
            return "0"
        sign = ""
        if num < 0:
            num *= -1
            sign = "-"
        res = ""
        while num > 0:
            res = str(num%7) + res
            num //= 7
        return sign + res
