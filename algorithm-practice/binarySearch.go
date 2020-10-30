package main

import "fmt"

func main() {
	// 前提是排好序的array
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(binarySearch(arr, 3))
}

func binarySearch(arr []int, n int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (left + right) / 2
		if arr[middle] > n {
			right = middle - 1
		} else if arr[middle] < n {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
