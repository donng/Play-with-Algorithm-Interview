package main

import "fmt"

func MinSubArrayLen(s int, nums []int) int {
	if s < 0 || nums == nil {
		panic("Illegal Arguments")
	}

	sum := 0
	l, r := 0, -1 // nums[l...r]为我们的滑动窗口

	length := len(nums)
	count := length + 1

	for l < length { // 窗口的左边界在数组范围内,则循环继续
		if r+1 < length && sum < s {
			r++
			sum += nums[r]
		} else { // r已经到头 或者 sum >= s
			sum -= nums[l]
			l++
		}

		if sum >= s {
			if (r - l + 1) < count {
				count = r - l + 1
			}
		}
	}

	if count == length+1 {
		return 0
	}

	return count
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	s := 15

	fmt.Println(MinSubArrayLen(s, nums))
}
