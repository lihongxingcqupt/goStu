package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported op: %s", op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("calling function %s with args "+"(%d, %d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	var sum = 0
	for i := range numbers {
		fmt.Println(i)
		sum += numbers[i]
	}
	return sum
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	i, err := eval(10, 20, "^")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	q, r := div(10, 2)
	fmt.Println(q, r)
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(sum(1, 1, 1, 1, 1))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
