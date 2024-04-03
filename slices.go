package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[2:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:6])
	fmt.Println(arr[:])

	fmt.Println("更新slice 2:6后")
	updateSlice(arr[2:6])
	fmt.Println(arr[2:6])
	fmt.Println(arr)

	fmt.Println("更新slice ：后")
	updateSlice(arr[:])
	fmt.Println(arr)
}
