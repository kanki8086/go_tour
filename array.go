package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := make([]int, 5)
	fmt.Println(len(a))
	// arrayInit()
	// app()
	app2()
}

func app2() {
	var s = make([]string, 0, 10)
	s = append(s, "tag_name=hoge")
	s = append(s, "tag_value=vale")
	fmt.Println(strings.Join(s, "&"))
}

func app() {
	primes := [5]int{2, 3, 5, 7, 11}
	newP := append(primes[0:], 13)
	fmt.Println(newP)
}

func arrayInit() {
	primes := [5]int{2, 3, 5, 7, 11}
	for i := 0; i < len(primes); i++ {
		fmt.Println(primes[i])
	}
	fmt.Println(primes[2:])
	fmt.Println(primes[1:3])

	primes[1:3][0] = 123
	fmt.Println(primes)
	fmt.Println("len=" + strconv.Itoa(len(primes)))
	fmt.Println("cap=" + strconv.Itoa(cap(primes)))
}

func arrayStr() {
	var ar [2]string
	ar[0] = "Hello "
	ar[1] = "world!"

	for i := 0; i < len(ar); i++ {
		fmt.Println(ar[i])
	}
}
