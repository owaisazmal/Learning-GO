package main

import "fmt"
//Go sees arrays as values
func main() {

	a := [...]int{1,2,3}
	b := &a		//b is assigned to copy of an a
	b[1] = 9
	fmt.Println(a)
	fmt.Println(b)

}