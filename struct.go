package main

import "fmt"
 
func main(){	

	aDoctor := struct{name string}{name: "John Peter"}		//structs are value types
	anotherDoctor := &aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}