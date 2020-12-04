package main

import (
	"fmt"
)

// ===
// 希尔排序是在直接插入排序的基础上改造而来的，
// 也就是说它的最后一次排序是 1也就变成了直接
// 插入排序，一般来说增量的时候是数组的长度/2
// ===

func main() {
	list := []int{1, 67, 33, 20, 5, 6, 71, 8, 0}
	fmt.Println(shellSort(list))

}

func shellSort(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}
	for step := len(arr) / 2; step >= 1; step /= 2 {
		fmt.Println("step===>", step)
		for i := step; i < len(arr); i += step {
			fmt.Println("i======>", i)
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				fmt.Println(arr[j], arr[j+step], i, j)
				arr[j], arr[j+step] = arr[j+step], arr[j]
				fmt.Println(arr[j], arr[j+step])
				fmt.Println(arr)
			}
		}
	}
	return arr
}
func demo1(arr []int) []int {
	for step := len(arr) / 2; step >= 1; step /= 2 {
		for i := step; i >= len(arr); i += step {
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]

			}
		}
	}
	return arr
}

func demo2(arr []int) []int {
	for step := len(arr) / 2; step >= 1; step /= 2 {
		for i := step; i < len(arr); i += step {
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]

			}
		}
	}
	return arr
}

func demo3(arr []int) []int {
	for step := len(arr) / 2; step >= 1; step /= 2 {
		for i := step; i < len(arr); i += step {
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}
	return arr
}

func demo4(arr []int) []int {
	for step := len(arr) / 2; step >= 1; step /= 2 {
		for i := step; i < len(arr); i += step {
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}
	return arr
}

func demo5(arr []int) []int {
	for step := len(arr) / 2; step >= 1; step /= 2 {
		for i := step; i < len(arr); i += step {
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}
	return arr
}

func demo6(arr []int) []int {
	for step := len(arr) / 2; step >= 1; step /= 2 {
		for i := step; i < len(arr); i += step {
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; {

			}
		}
	}
	return arr
}

func demo7(arr []int) []int {
	for step := len(arr) / 2; step >= 1; step /= 2 {
		for i := step; i < len(arr); i += step {
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}
	return arr
}
