package main

import "fmt"

// 167. Two Sum II - Input array is sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
//
// 二分搜索法
// 时间复杂度: O(nlogn)
// 空间复杂度: O(1)
func TwoSum(numbers []int, target int) []int {
	length := len(numbers)

	if length < 2 {
		panic("Illegal argument numbers")
	}

	for i := 0; i < length; i++ {
		// 二分查找
		j := binarySearch(numbers, i+1, length-1, target-numbers[i])
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}

	panic("The input has no solution")
}

func binarySearch(nums []int, l int, r int, target int) int {
	if l < 0 || l > len(nums) {
		panic("l is out of bound")
	}

	if r < 0 || r > len(nums) {
		panic("r is out of bound")
	}

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(TwoSum(nums, target))
}
