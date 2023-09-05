package main

import (
	"fmt"
	"strconv"
)

func main() {	//overloading/converting
	
	var i float32 = 42.5
	fmt.Printf("%v, %T\n", i, i)

	var j int 
	j = int(i)			// Note: can't type j=i (have to type explixitly)
	fmt.Printf("%v, %T\n", j, j)

	var k string 
	k = strconv.Itoa(j)			// Note: can't type k = string(j) 		
	fmt.Printf("%v, %T\n", k, k)

}