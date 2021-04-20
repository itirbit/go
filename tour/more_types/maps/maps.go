package main

import "fmt"
import "golang.org/x/tour/wc"

func maps(){

	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func mutate_maps(){
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int{
	var m = make(map[string]int)
	var buff = ""
	for i:=0; i <len(s); i++{
		if s[i] == ' '{
			m[buff]++
			buff = ""
		}else{
			buff += string(s[i])
		}
	}
	m[buff]++
	return m
}

func maps_exe(){
	wc.Test(WordCount)
}

func main(){
	maps_exe();
}
