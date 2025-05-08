package main

/*
- Khi gặp một số  => tạm lưu lại vào num.
- Khi gặp một toán tử mới (như +, -, *, /):
    - Dùng sign để xử lý số num vừa đọc xong.
			- pre sign: +, - thì append vào stack
			- pre sign: *, / thì cần tính giá trị cuối cùng của stack
    - Gán sign = toán tử mới.

* Nếu vị trí cuối cùng của chuỗi là số hoặc khoảng trắng ' ' 
(nghĩa là sau số cuối không có thêm toán tử nào), thì bạn phải thực hiện xử lý số cuối cùng với sign hiện tại.
*/
// TC: O(n), SC: O(n)
func calculate(s string) int {
	nums := []int{}
	num := 0
	sign := byte('+')

	for i := 0; i < len(s); i++ {
		c := s[i]
		if isDigit(c) {
			num = num*10 + int(c-'0')
		}

		if (c != ' ' && !isDigit(c)) || i == len(s)-1 {
			switch sign {
			case '+':
				nums = append(nums, num)
			case '-':
				nums = append(nums, -num)
			case '*':
				nums[len(nums)-1] *= num
			case '/':
				nums[len(nums)-1] /= num
			}

			num = 0
			sign = s[i]
		}

	}
	res := 0
	for _, num := range nums {
		res += num
	}

	return res

}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
