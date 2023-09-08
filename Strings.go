//String concatination
package main

import "fmt"

func main(){
	s := "this is a string"
	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s + s2, s + s2)
	
	//printing ascii value

	x := "this is a string"
	y := []byte(x)
	fmt.Printf("%v, %T\n", y, y)

	//rune ***imp
	//rune in Go is a data type that stores codes that represent Unicode characters. 
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}