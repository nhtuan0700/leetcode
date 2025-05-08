package main

// https://leetcode.com/problems/beautiful-array/

// Intuition:
// - Nếu ta chia mảng thành hai nửa tách biệt là các số lẻ và các số chẵn,
//   thì không thể có giá trị nào ở giữa hai số đó tạo thành trung bình.
// → Vì: trung bình của hai số chẵn hoặc hai số lẻ là chẵn,
//       còn trung bình của một chẵn và một lẻ là lẻ.
//   Nhưng nếu chẵn và lẻ nằm ở hai nửa tách biệt thì không thể tạo thành bộ ba hợp lệ.

// Tính chất bảo toàn:
// - Với mọi mảng A là beautiful array:
//   - Cộng:
//     Nếu 2*A[k] ≠ A[i] + A[j], thì:
//     2*(A[k] + x) = 2*A[k] + 2x ≠ A[i] + A[j] + 2x = (A[i]+x) + (A[j]+x)
//   - Nhân:
//     2*(A[k]*x) = 2*A[k]*x ≠ (A[i]+A[j])*x = A[i]*x + A[j]*x
// → Biến đổi tuyến tính f(x) = ax + b (với a ≠ 0) giữ nguyên tính chất beautiful.

// Ý tưởng thuật toán:
// - Bắt đầu với mảng beautiful nhỏ nhất: [1]
// - Dùng biến đổi tuyến tính để sinh mảng mới:
//     - A1 = [2*x - 1 for x in A] → toàn bộ là số lẻ
//     - A2 = [2*x     for x in A] → toàn bộ là số chẵn
// - Kết hợp A1 + A2 để tạo beautiful array lớn hơn

// Mẹo:
// Tưởng tượng tách mảng làm 2 nửa không thể "trung bình hóa" nhau.
// Luôn sinh ra số lẻ trước rồi số chẵn sau từ các giá trị hiện có.
// Dựa vào tính chất đệ quy: nếu res là beautiful array, thì [2*x-1 for x in res] + [2*x for x in res] cũng là beautiful array.
// TC: O(n), SC: O(n)
func beautifulArray(n int) []int {
	res := []int{1}

	for len(res) < n {
		tmp := []int{}

		// event part
		for _, num := range res {
			even := num * 2
			if even <= n {
				tmp = append(tmp, even)
			}
		}
		// odd part
		for _, num := range res {
			odd := num*2 - 1
			if odd <= n {
				tmp = append(tmp, odd)
			}
		}

		res = tmp
	}

	return res
}
