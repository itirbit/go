package main

import "fmt"
import "math"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Absp() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func indirection1(){
	v := Vertex{3,4}
	v.Scale(2)
	ScaleFunc(&v,10)

	p := &Vertex{4,3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v,p)
}

func indirection2(){
	v := Vertex{3,4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p:= &Vertex{4,3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

func pointer_receiver(){
	v := &Vertex{3,4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

func pointer_receiver2(){
	v := &Vertex{3,4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Absp())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Absp())
}

func main(){
	pointer_receiver2()
}
