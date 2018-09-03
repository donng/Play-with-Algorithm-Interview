package Util

import "math/rand"

func GenerateRandomArray(n int, rangeL int, rangeR int) []int {
	if n <= 0 || rangeL > rangeR {
		panic("invalid params")
	}

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Int()%(rangeR-rangeL+1) + rangeL
	}

	return arr
}

func GenerateOrderedArray(n int) []int {
	if n <= 0 {
		panic("invalid params")
	}

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}

	return arr
}
