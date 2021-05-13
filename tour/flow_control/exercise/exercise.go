package main

import "fmt"

func Sqrt(x int) float64 {
	z := 1.0
	for i:=1; i<10;i++{
		z -= (z*z-float64(x))/(2*z)
	}
	return z
}

func main(){
	i := 1
	for i<100{
		fmt.Printf("value: %v, sqrt:%v\n", i, Sqrt(i))
		i += i
	}
}
