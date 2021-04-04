package main

import "fmt"

type Vertex struct {
	x,y int
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{x: 1}
	v3 = Vertex{}
	p = &Vertex{1,2}
)

func main() {
	p.x = 2
	p.y = 4
	v1.x = 4
	v2.y = 8
	fmt.Println(v1, p, v2, v3)
}
