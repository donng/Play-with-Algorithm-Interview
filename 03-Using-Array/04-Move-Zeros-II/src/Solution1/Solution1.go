package main

import "fmt"

// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/description/
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func MoveZeroes(nums []int) {

	var nonZeroElements []int

	// 将vec中所有非0元素放入nonZeroElements中
	for _, value := range nums {
		if value != 0 {
			nonZeroElements = append(nonZeroElements, value)
		}
	}

	// 将nonZeroElements中的所有元素依次放入到nums开始的位置
	for index, value := range nonZeroElements {
		nums[index] = value
	}

	// 将nums剩余的位置放置为0
	for i := len(nonZeroElements); i < len(nums); i++ {
		nums[i] = 0
	}
}

func main()  {
	arr := []int{0, 1, 0, 3, 12}

	MoveZeroes(arr)

	fmt.Println(arr)
}