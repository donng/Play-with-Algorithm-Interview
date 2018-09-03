package BinarySearch

func BinarySearch(arr []int, n int, target int) int {
	l, r := 0, n // 在 [l...r) 的范围里寻找 target
	for l < r {  // 当 l == r时,区间 [l...r) 是无效区间
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			r = mid // target在[l...mid)中; [mid...r]一定没有target
		} else {    // target > arr[mid]
			l = mid + 1 // target在[mid+1...r)中; [l...mid]一定没有target
		}
	}

	return -1
}
