# https://leetcode.com/problems/online-stock-span/

# TC: O(n), SC: O(n)
class StockSpanner:
    def __init__(self):
            self.stack = [] # (price, span)

    def next(self, price: int) -> int:
        span = 1
        while self.stack and self.stack[-1][0] <= price:
            span += self.stack.pop()[1]

        self.stack.append((price, span))
        return span

# use monotonic stack
# self.stack lưu index của các phần tử có thể là đỉnh cao nhất trước đó (theo quy tắc monotonic decreasing stack).
# Khi thêm price mới:
# - Bạn pop các giá nhỏ hơn hoặc bằng để giữ stack giảm dần.
# span được tính là:
# - Nếu stack rỗng → tất cả giá trước đó đều nhỏ hơn hoặc bằng → span = len(prices) + 1
# - Nếu không rỗng → span = len(prices) - stack[-1]
class StockSpanner:
    def __init__(self):
        self.stack = []
        self.prices = []

    def next(self, price: int) -> int:
        while self.stack and self.prices[self.stack[-1]] <= price:
            self.stack.pop()

        span = 1
        if not self.stack:
            span = len(self.prices) + 1
        else:
            span = len(self.prices) - self.stack[-1]
            
        self.stack.append(len(self.prices))
        self.prices.append(price)
        return span


# brute force
# TC: O(n^2), SC: O(n)
class StockSpanner:
    def __init__(self):
        self.prices = []

    def next(self, price: int) -> int:
        j = len(self.prices)
        for i in range(len(self.prices) -1, -1, -1):
            if self.prices[i] > price:
                break
            j = i
        self.prices.append(price)
        return len(self.prices) - j


# Your StockSpanner object will be instantiated and called as such:
# obj = StockSpanner()
# param_1 = obj.next(price)
