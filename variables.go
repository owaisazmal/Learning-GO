package main

import (
	"fmt"
)

// package level declaration
// Variables "Always" needed to be used
var x float32 = 4

func main() { // 3 Ways to declare

	var i float32 = 1 //i
	fmt.Println(i)

	var j int //ii
	j = 2
	j = 69
	fmt.Println(j)
	//fmt.Println(423)

	k := 3 //it can figure the data type by itself				iii
	fmt.Println(k)

	// Formating String
	fmt.Printf("%v, %T", x, x) // %v->the var, %T->type

}
