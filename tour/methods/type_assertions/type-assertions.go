package main

import "fmt"

func main(){
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	if (ok) {
		fmt.Println("cast to string success ", s)
	}else{
		fmt.Println("failed cast to string")
	}

	f, ok := i.(float64)
	if (ok){
		fmt.Println("cast to float64 success ", f)
	}else{
		fmt.Println("failed cast to float64")
	}

	in, ok := i.(int)
	if(ok){
		fmt.Println("cast to int success ", in)
	}else{
		fmt.Println("failed cast to int")
	}
}


