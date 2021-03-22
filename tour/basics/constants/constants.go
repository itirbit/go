package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "Dafuq"
	World = "change world"
	fmt.Println("hello", World)
	Pi = 3.14159
	fmt.Println("happy", Pi, "Day")

	const Truth = true
	Truth = false
	fmt.Println("Go rules?", Truth)
}
