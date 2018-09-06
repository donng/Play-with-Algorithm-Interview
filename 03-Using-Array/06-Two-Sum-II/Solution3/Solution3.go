package main

import "fmt"

// 167. Two Sum II - Input array is sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
//
// 对撞指针
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func TwoSum(numbers []int, target int) []int {
	length := len(numbers)
	if length < 2 {
		panic("Illegal argument numbers")
	}

	l := 0
	r := length - 1

	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}

	panic("The input has no solution")
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(TwoSum(nums, target))
}
