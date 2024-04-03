package main

import "fmt"

func printArr(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	//
	fmt.Println(arr1, arr2, arr3)

	sum := 0
	for _, v := range arr2 {
		sum += v
	}
	fmt.Println(sum)

	fmt.Println("调用方法打印")
	printArr(&arr3)
	fmt.Println("直接打印")
	fmt.Println(arr3)
}
