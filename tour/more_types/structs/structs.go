package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1,2})
	v := Vertex{1,2}
	v.x = 4
	fmt.Println(v.y)
}
