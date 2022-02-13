package main

import (
	"fmt"
)

type Employee struct {
	id   int
	role string
}

func main() {
	m := make(map[string]Employee)
	m["alice"] = Employee{1, "accounting"}
	m["bob"] = Employee{2, "engineer"}
	fmt.Println(m)
	fmt.Println(m["alice"])
	fmt.Println(m["none"])
}
