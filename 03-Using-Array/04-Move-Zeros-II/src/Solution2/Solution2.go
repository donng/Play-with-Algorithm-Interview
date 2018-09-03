package main

import "fmt"

// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/description/
//
// 原地(in place)解决该问题
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func MoveZeroes(nums []int)  {
	// nums中, [0...k)的元素均为非0元素
	k := 0

	// 遍历到第i个元素后,保证[0...i]中所有非0元素
	// 都按照顺序排列在[0...k)中
	for _, value := range nums {
		if value != 0 {
			nums[k] = value
			k++
		}
	}

	// 将nums剩余的位置放置为0
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 12}

	MoveZeroes(arr)

	fmt.Println(arr)
}
