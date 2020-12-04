package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12, 87, 66}
	fmt.Println(recursionSort(arr))
}

// ===
// 递归快速排序 （缺点是太耗费内存）
// ===
func recursionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	split := arr[0]
	low := make([]int, 0, 0)
	high := make([]int, 0, 0)
	middle := make([]int, 0, 0)
	// 把选出来的第一个数字先加进去
	middle = append(middle, split)
	for i := 1; i < len(arr); i++ {
		if arr[i] < split {
			low = append(low, arr[i])
		} else if arr[i] > split {
			high = append(high, arr[i])
		} else {
			middle = append(middle, arr[i])
		}
	}
	// 递归把两边排一下
	low, high = recursionSort(low), recursionSort(high)
	newArr := append(append(low, middle...), high...)
	return newArr
}
