package main

import "fmt"

func main() {

	var i interface{} = true

	switch i.(type){
	case int:
		fmt.Println("It's an integer")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("I don't know what 'i' is ")
	}
}
