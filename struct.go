package main

import "fmt"
// Structs: Collections of disparate data types that describe a single concept
//Keyed by names fields
//structs are value type
func main(){	

	aDoctor := struct{name string}{name: "John Peter"}		//structs are value types
	anotherDoctor := &aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}