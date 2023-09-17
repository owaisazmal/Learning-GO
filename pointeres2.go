package main

import "fmt"

func main(){
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)		//&myStruct{foo: 50}		//same thing
	(*ms).foo = 50
	fmt.Println(ms)
}

type myStruct struct{
	foo int
}