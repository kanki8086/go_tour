package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

func modifyVertex(v *Vertex) {
	v.x = 7
	// (*v).x = 6
}

func main() {
	var v = Vertex{1, 2}
	modifyVertex(&v)
	fmt.Println(v)
}
