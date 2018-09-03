package main

import (
	"math"
	"time"
	"fmt"
	"practice/Play-with-Algorithm-Interview/03-Using-Array/01-Binary-Search/src/Util"
	"practice/Play-with-Algorithm-Interview/03-Using-Array/01-Binary-Search/src/BinarySearch"
)

func main() {
	n := int(math.Pow(10, 7))
	data := Util.GenerateOrderedArray(n)

	startTime := time.Now()
	for i := 0; i < n; i++ {
		if i != BinarySearch.BinarySearch(data, n, i) {
			panic("find i failed!")
		}
	}

	fmt.Println("Binary Search test complete.")
	fmt.Println("Time cost: ", time.Now().Sub(startTime))
}
