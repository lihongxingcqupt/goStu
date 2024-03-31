package main

import (
	"fmt"
	"math"
)

func getTriangle(x int, y int) int {
	var a int = int(math.Sqrt(float64(x*x + y*y)))
	return a
}

func main() {
	fmt.Println("hello world")
	fmt.Println(getTriangle(3, 4))
}
