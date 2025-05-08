# https://leetcode.com/problems/string-to-integer-atoi/

# TC: O(n), SC: O(1)
class Solution:
	def myAtoi(self, s: str) -> int:
		startIdx = 0
		signed = 0
		for c in s:
			if c == ' ':
				startIdx += 1
				continue
			if c == '+' or c == '-':
				signed = 1 if c == '+' else -1
				startIdx += 1
				break
			break
		res = 0
		MAX_INT = 2**31 - 1
		MIN_INT = -2**31
		while startIdx < len(s):
			if not s[startIdx].isnumeric():
				break
			num = ord(s[startIdx]) - ord('0')
			# -2^31 <= res * 10 + num <= 2^31 - 1
			# (-2^31 - num) / 10 <= res <= (2^31 - 1 - num) / 10
			# if res > (MAX_INT - num) / 10:
			#     if signed == -1:
			#         res = MAX_INT + 1
			#     else:
			#         res = MAX_INT
			#	  break 
			# res = res * 10 + num
			if signed == -1:
				if -res < int((MIN_INT + num) / 10):
					return MIN_INT
			else:
				if res > int((MAX_INT - num) / 10):
					return MAX_INT
			res = res * 10 + num
			startIdx += 1
		
		res = res * signed if signed != 0 else res
		return res
		
		