package main

import "fmt"
//creating a pointer
// a pointer is a variable that stores the address of an object stored in memory. 
func main(){

	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	a = 43
	fmt.Println(&a, *b)			// -> & for getting the address (memory location), * for getting the value of the address
}