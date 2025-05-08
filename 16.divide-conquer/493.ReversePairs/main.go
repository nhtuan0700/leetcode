package main

// Có 3 trường hợp cặp (i, j) trong nums:
// 1. Cả i và j đều nằm trong đoạn trái (left part)
// 2. Cả i và j đều nằm trong đoạn phải (right part)
// 3. i nằm ở đoạn trái, j nằm ở đoạn phải
// - Với trường hợp 1 và 2:
//   - Các cặp này đã được đếm ở các bước đệ quy con trước đó (tức là gọi mergeSort(l, m) và mergeSort(m+1, r))
//   - Vì thế ta chỉ cần cộng lại kết quả từ 2 bên
// - Với trường hợp 3 (i ở trái, j ở phải):
//   - Đây là phần quan trọng cần xử lý tại bước merge hiện tại
//   - Ta đếm số lượng j sao cho nums[i] > 2 * nums[j]
//   - Vì 2 nửa đã được sort, ta có thể dùng 2 con trỏ (two pointers) để tối ưu quá trình đếm
// - Sau khi đếm xong:
//   - Cần merge 2 nửa [l..m] và [m+1..r] lại thành mảng đã sort
//   - Việc này giúp đảm bảo các bước merge cao hơn trong đệ quy vẫn có thể áp dụng two pointers

// https://leetcode.com/problems/reverse-pairs/
// merge sort
// TC: O(nlogn), SC: O(n)
func reversePairs(nums []int) int {
	count := mergeSort(nums, 0, len(nums)-1)
	return count
}

func mergeSort(nums []int, l, r int) int {
	if l == r {
		return 0
	}

	m := l + (r-l)/2
	count := mergeSort(nums, l, m) + mergeSort(nums, m+1, r)

	j := m + 1
	for i := l; i <= m; i++ {
		for j <= r && nums[i] > 2*nums[j] {
			j++
		}
		// m+1 ... j-1
		count += j - (m + 1)
	}

	merge(nums, l, m, r)
	return count
}

func merge(nums []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	L := make([]int, n1)
	R := make([]int, n2)
	for i := 0; i < n1; i++ {
		L[i] = nums[i+l]
	}
	for i := 0; i < n2; i++ {
		R[i] = nums[i+m+1]
	}

	i, j := 0, 0
	k := l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			nums[k] = L[i]
			i++
		} else {
			nums[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		nums[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		nums[k] = R[j]
		j++
		k++
	}
}
