package main

import "fmt"

// 75. Sort Colors
// https://leetcode.com/problems/sort-colors/description/
//
// 计数排序的思路
// 对整个数组遍历了两遍
// 时间复杂度: O(n)
// 空间复杂度: O(k), k为元素的取值范围
func SortColors(nums []int) {
	count := make([]int, 3)
	for _, v := range nums {
		if v < 0 || v > 2 {
			panic("Invalid params.")
		}
		count[v]++
	}

	index := 0
	for i, v := range count {
		for j := 0; j < v; j++ {
			nums[index] = i
			index++
		}
	}
	//for i := 0; i < count[0]; i++ {
	//	nums[index] = 0
	//	index++
	//}
	//for i := 0; i < count[1]; i++ {
	//	nums[index] = 1
	//	index++
	//}
	//for i := 0; i < count[2]; i++ {
	//	nums[index] = 2
	//	index++
	//}
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	SortColors(arr)

	fmt.Println(arr)
}
