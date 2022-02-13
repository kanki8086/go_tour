package main

import (
	"fmt"
)

func ret() string {
	fmt.Println("start!")
	return "world"
}

func main() {
	defer fmt.Println(ret())
	fmt.Println("main hello")
}
