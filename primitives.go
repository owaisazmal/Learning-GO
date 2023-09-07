//Working with boolean
package main

import "fmt"

func main() {
	// var n bool = true

	n := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", m, n)
	fmt.Printf("%v, %T\n", n, n)

	var x bool						//0(false) value for boolean by default
	fmt.Printf("%v, %T\n", x, x)

	var a int  = 10
	var b int8 = 3
	fmt.Println(a + int(b))			//can't type fmt.Println(a + b)  you will have to convert it (typecast)
}