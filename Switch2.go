package main

import "fmt"

func main(){
	var i interface{} = [3]int{}
	switch i.(type) {
	case int: 
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is an String")
	case [3]int:
		fmt.Println("i is an [3] int")
	default:
		fmt.Println("i is an another type")
	}
}