package main

import "fmt"

func main() {

	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12, 87, 66}

	fmt.Println(insertSort(arr))
}

// 直接插入快排序 （缺点是一亿数据就得循环一亿次）
func insertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		// arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12, 87, 66}
		for i := 1; i < len(arr); i++ {
			temp := arr[i]
			j := i - 1
			for j >= 0 && temp < arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			// j = -1
			arr[j+1] = temp
		}
		return arr
	}
}
