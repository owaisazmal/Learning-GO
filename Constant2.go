//9/7/23

package main

import (
	"fmt"
	//"math"
)

//Numerated Constant
const (
	a = iota				//important (Iota in Go is used to represent constant increasing sequences.)
	b = iota
	c = iota
)

const (
	a2 = iota				//DEPICTS AND RESETS
	b2 
	c2 
)
func main(){

	const myConst = 4									//it can determine the value
	var x int16 = 27
	fmt.Printf("%v, %T\n", myConst+x, myConst+x)

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)

	fmt.Printf("%v, %T\n", a2, a2)
	fmt.Printf("%v, %T\n", b2, b2)
	fmt.Printf("%v, %T\n", c2, c2)


}