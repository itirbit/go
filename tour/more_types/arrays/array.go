package main

import "fmt"

func arrays(){
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}

func slices(){
	primes := [6]int{2,3,5,7,11,13}
	var s[]int = primes[1:4]
	fmt.Println(s)
}

func slices_pointers(){
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a:=names[0:2]
	b:=names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a,b)
	fmt.Println(names)
}

func main(){
	slices_pointers()
}
