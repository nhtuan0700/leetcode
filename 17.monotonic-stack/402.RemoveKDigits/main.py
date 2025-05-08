from collections import deque

# idea:  To make the number as small as possible, remove digits from the stack if the current digit is smaller than the top of the stack and k > 0.

# https://leetcode.com/problems/remove-k-digits/
# TC: O(n)  SC: O(n - k)
class Solution1:
    def removeKdigits(self, num: str, k: int) -> str:
        if k == len(num):
            return "0"

        res = [] # monotonic stack
        for i in range(len(num)):
            while res and int(res[-1]) > int(num[i]) and k > 0:
                res.pop()
                k -= 1
            res.append(num[i])

        while k > 0:
            res.pop()
            k -= 1

        start = 0
        while start + 1 < len(res) and res[start] == "0":
            start += 1

        return "".join(res[start:])

        
# find min in range queries => using monotonic queue
# find min in (i, i+k)

# ✅ Ý tưởng:
# Ta cần giữ lại n - k chữ số.
# Với mỗi vị trí i trong 0 → n - k - 1, ta tìm chữ số nhỏ nhất trong đoạn [i, i + k].
# Vì một khi đã chọn chữ số tại vị trí j, thì các chữ số trước j đều bị coi là đã bị xóa → đúng logic của bài toán.

# 📌 Giải thích chi tiết:
# Sử dụng monotonic queue tăng dần để giữ chỉ số của các chữ số ứng viên cho min.
# Với mỗi bước i, đảm bảo:
# - Queue chứa chỉ số trong phạm vi [i, i + k]
# - Đầu queue là chỉ số chữ số nhỏ nhất → chọn nó.
# j di chuyển để thêm vào queue và đảm bảo mỗi chữ số được xét đúng một lần → O(n).

# "1432219", k = 3
# output _ _ _ _ (n - k = 7 - 3 = 4)
# i = 0 -> range(0,3) -> min = 1 (pos: 0)
# i = 1 -> range(1,4) -> min = 2 (pos: 3)
# i = 2 -> range(4,5) -> min = 1 (pos: 5)
# i = 3 -> range(6,6) -> min = 9 (pos: 6)

# i -> n - k
#   j: range(?,i+k)
# TC: O(n) SC: O(n-k)
class Solution2:
    def removeKdigits(self, num: str, k: int) -> str:
        n = len(num)
        if k == len(num):
            return "0"
        dq = deque([])
        j = 0
        res = ""
        for i in range(n-k):
            while j <= i + k:
                while dq and num[dq[-1]] > num[j]:
                    dq.pop()
                dq.append(j)
                j += 1
            res += num[dq[0]]
            dq.popleft()
 
        start = 0
        while start + 1 < len(res) and res[start] == "0":
            start += 1

        return res[start:]
