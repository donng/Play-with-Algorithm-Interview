package main

import "fmt"

// 209. Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/description/
//
// 优化暴力解
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func MinSubArrayLen(s int, nums []int) int {
	if s < 0 || nums == nil {
		panic("Illegal Arguments")
	}

	length := len(nums)
	count, minimal := 0, length+1

	// sums[i]存放nums[0...i-1]的和
	sums := make([]int, length + 1)
	sums[0] = 0
	for i := 1; i <= length; i++ {
		sums[i] = nums[i-1] + sums[i-1]
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			// 使用sums[j+1] - sums[i] 快速获得nums[i...j]的和
			sum := sums[j+1] - sums[i]
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
