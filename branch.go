package main

import (
	"fmt"
	"io/ioutil"
)

func getGrade(grade int) string {
	g := ""
	switch {
	case grade > 100 || grade < 0:
		panic(fmt.Sprintf("成绩有误 grade = %d", grade))
	case grade < 60:
		g = "F"
	case grade < 80:
		g = "C"
	case grade <= 100:
		g = "A"
	}
	return g
}
func main() {
	if file, err := ioutil.ReadFile("abc.txt"); err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s\n", file)
	}

	fmt.Println(
		getGrade(100),
		getGrade(90),
		getGrade(50),
		getGrade(70))
}
