package main

import "fmt"

func main(){
	a := 10													//1010
	b := 3													//0011
	fmt.Println(a & b)	//and								//0010 -> 2 
	fmt.Println(a | b)	//or								//1011 -> 11
	fmt.Println(a ^ b)	//exclusive or						//1001 -> 9
	fmt.Println(a &^ b)	//and not							//0100 -> 8

}