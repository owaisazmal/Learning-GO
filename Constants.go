/*Agenda
1.Naming Conventions
2.Typed Constants
3.Untyped Constants
4.Enumerated Constants
5.Enumerated Expression
*/
package main

import (
	"fmt"
	//"math"
)
const myConst int = 4			//can be shaddow like variables (inner variable will always win)
func main(){

	const myConst int = 42 			//Namin covention
	// myCont = 27 					//not allowed to change the value like int,etc
	fmt.Printf("%v, %T\n", myConst, myConst)

	//const myConst2 float64 = math.Sin(1.57)			//cant do this
	// fmt.Printf("%v, %T\n", myConst2, myConst2)

	 const a int = 14
	 const b string = "Pizza"
	 const c float32 = 3.14
	 const d bool = true
	 fmt.Printf("%v, %T\n", a)
	 fmt.Printf("%v, %T\n", b)
	 fmt.Printf("%v, %T\n", c)
	 fmt.Printf("%v, %T\n", d)

}