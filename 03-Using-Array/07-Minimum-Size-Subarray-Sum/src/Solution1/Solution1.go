package main

import "fmt"

// 209. Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/description/
//
// 暴力解法
// 该方法在 Leetcode 中会超时！
// 时间复杂度: O(n^3)
// 空间复杂度: O(1)
func MinSubArrayLen(s int, nums []int) int {
	if s < 0 || nums == nil {
		panic("Illegal Arguments")
	}

	length := len(nums)
	count, minimal := 0, length+1
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			count = j - i + 1
			if sum >= s && count < minimal {
				minimal = count
			}
		}
	}

	if minimal == length+1 {
		return 0
	}

	return minimal
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	s := 15

	fmt.Println(MinSubArrayLen(s, nums))
}
