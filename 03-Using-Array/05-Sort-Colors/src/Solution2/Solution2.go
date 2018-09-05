package main

import "fmt"

// 75. Sort Colors
// https://leetcode.com/problems/sort-colors/description/
//
// 三路快速排序的思想
// 对整个数组只遍历了一遍
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func SortColors(nums []int) {
	l := -1         // [0...zero] == 0
	r := len(nums)  // [two...n-1] == 2

	for i := 0; i < r; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 0 {
			l++
			nums[l], nums[i] = nums[i], nums[l]
			i++
		} else {
			if nums[i] !=2 {
				panic("Invalid params")
			}
			r--
			nums[r], nums[i] = nums[i], nums[r]
		}
	}
}

func main() {
 	arr := []int{2, 0, 1}
	SortColors(arr)

	fmt.Println(arr)
}
