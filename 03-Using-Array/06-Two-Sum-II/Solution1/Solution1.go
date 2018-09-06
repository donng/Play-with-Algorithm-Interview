package main

import "fmt"

// 167. Two Sum II - Input array is sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
//
// 暴力枚举法
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func TwoSum(numbers []int, target int) []int {
	length := len(numbers)

	if length < 2 {
		panic("Illegal argument numbers")
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if numbers[i] + numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	panic("The input has no solution")
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(TwoSum(numbers, target))
}
