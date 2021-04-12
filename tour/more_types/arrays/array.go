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

func slice_literals(){
	q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	r := []bool{true,false,true,true,false,true}
	fmt.Println(r)

	s:=[]struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func slice_bounds() {
	s := []int{2,3,5,7,11,13}

	s = s[:len(s)-1]
	fmt.Println(s)

	s = s[:len(s)-1]
	fmt.Println(s)

	s = s[:len(s)-1]
	fmt.Println(s)

	s = s[:len(s)-1]
	fmt.Println(s)

	s = s[:len(s)-1]
	fmt.Println(s)
}

func slice_len_cap(){
	s:=[]int{2,3,5,7,11,13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nil_slices(){
	var s[]int
	fmt.Println(s, len(s), cap(s))
	if s == nil{
		fmt.Println("nil!")
	}
}

func making_slices(){
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c:= b[:2]
	printSlice(c)

	d:=c[2:5]
	printSlice(d)
}

func main(){
	making_slices()
}
