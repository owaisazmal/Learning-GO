package main

import "fmt"

func main(){
	a := [3]int{1,2,3}
	b := &a[0]
	c := &a[1] 				//-4 can't do math on pointers
	fmt.Printf("%v %p %p\n", a , b, c)
}