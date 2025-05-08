package main

// https://leetcode.com/problems/count-of-range-sum/

// sum(i, j) = sum[j] - sum[i-1]
// Nếu i đứng trước phần tử đầu tiên thỏa mãn ta có thể viết lại
// sum(i, j) = sum[j] - sum[i]
// bài toán yêu cầu: lower <= sum[j] - sum[i] <= upper
// - điều kiện: 0 <= i < j <= n
// test case: [1,2,-1,3,-5] expected: 9

// Khi ta xét sum[i], ta tìm khoảng (j_start, j_end) sao cho:
//      lower ≤ sums[j] - sums[i] ≤ upper
//   ⇔  sums[j] ∈ [sums[i] + lower, sums[i] + upper]
//
// Vì mảng sums[] đã được sắp xếp trong bước merge,
// khi sum[i] tăng, khoảng tìm kiếm [j_start, j_end] cũng sẽ dịch sang phải.
//
// Do đó, ta chỉ cần tăng j_start và j_end dần về phía trước,
// không cần quay lại phía sau vì các phần tử đó chắc chắn không còn thỏa mãn điều kiện.
//
// Việc này giúp giữ thời gian chạy của vòng lặp tổng thể là O(n),
// thay vì O(n²) nếu ta reset j_start mỗi lần.

// TC: O(nlogn), SC: O(n)
func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	sums := make([]int, n+1)

	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}

	return mergeSort(sums, lower, upper, 0, n)
}

func mergeSort(nums []int, lower, upper, l, r int) int {
	if l == r {
		return 0
	}

	m := l + (r-l)/2
	res := mergeSort(nums, lower, upper, l, m) + mergeSort(nums, lower, upper, m+1, r)

	jStart, jEnd := m+1, m+1
	for i := l; i <= m; i++ {
		for jStart <= r && nums[jStart]-nums[i] < lower {
			jStart++
		}

		for jEnd <= r && nums[jEnd]-nums[i] <= upper {
			jEnd++
		}
		// sum_j in (jStart ... jEnd - 1)
		res += jEnd - jStart
	}

	// merge and sort
	n := r - l + 1
	tmp := make([]int, n)
	k := 0
	for i, j := l, m+1; i <= m || j <= r; k++ {
		if j > r || (i <= m && nums[i] <= nums[j]) {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
	}
	for i := 0; i < n; i++ {
		nums[i+l] = tmp[i]
	}

	return res
}
