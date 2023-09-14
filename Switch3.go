package main

import "fmt"

func main(){
	var i interface{} = 1
	switch i.(type){
	case int: 
		fmt.Println("i is an int")
		break			//only print the above line
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is an String")
	default:
		fmt.Println("i is an another type")
	}
}