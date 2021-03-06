package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64{
	return fn(3,4)
}

func functions(){
	hypot := func(x,y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fib() func() int{
	a := 0
	b := 1
	return func() int {
		t := a
		a = a + b
		b = t
		return b
	}
}

func fibonacci(){
	f := fib()
	for i:=0; i < 10; i++ {
		fmt.Println(f())
	}
}

func function_closures(){
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func main() {
	fibonacci();
}
