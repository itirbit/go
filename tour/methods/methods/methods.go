package main

import (
	"fmt"
	"math"
)
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func func1(){
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func func2(){
	v := Vertex{3,4}
	fmt.Println(Abs(v))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f<0 {
		return float64(-f)
	}
	return float64(f)
}

func func3(){
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func main(){
	func1()
}
