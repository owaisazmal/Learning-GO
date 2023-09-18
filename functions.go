package main

import "fmt"
// Basic syntax
func main(){

	for i := 0; i < 5; i++{

		sayMsg("Hello Go!", i)
	}
}

func sayMsg(msg string, idx int){

	fmt.Println(msg)
	fmt.Println("This value of the index is ", idx)
}

