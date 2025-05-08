# https://leetcode.com/problems/split-a-string-in-balanced-strings/

# TC: O(n), SC: O(1)
class Solution:
    def balancedStringSplit(self, s: str) -> int:
        balance = 0
        res = 0
        for c in s:
            if c == 'R':
                balance += 1
            else:
                balance -= 1
            if balance == 0:
                res += 1
        
        return res

# Greedy: 
# - output: maximum number of splited balanced string
# Mỗi lần bạn duyệt từ trái sang phải, bạn ngay lập tức tách chuỗi khi gặp một chuỗi cân bằng.
# Bạn không cần nhìn xa, chỉ cần khi nào đạt balance == 0 thì tăng res lên 1.
# Đây là quyết định cục bộ tối ưu tại mỗi bước và cũng đảm bảo tối ưu toàn cục, vì nếu bạn không tách ra ngay tại lúc balance == 0, thì bạn sẽ trộn nó vào một chuỗi lớn hơn => làm giảm số lượng chuỗi con có thể chia.
