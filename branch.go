package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	if file, err := ioutil.ReadFile("abc.txt"); err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s\n", file)
	}
}
