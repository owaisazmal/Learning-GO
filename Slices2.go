package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   //SLice all the elements
	c := a[3:]  //SLice from 4th element to end
	d := a[:6]  //Slice first 6 elements
	e := a[3:6] //slice 4th, 5th and 6th elements
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
