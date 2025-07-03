package main

import (
	"fmt"
)

func main() {
	var n int
	var arr []int

	fmt.Scan(&n)
	arr = make([]int, n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		arr[i] = x
	}

	mergeSort(arr, 0, n-1)

	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
}
