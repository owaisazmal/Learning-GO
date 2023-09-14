package main

import "fmt"
// use fallthrough to remove break;
func main(){
	switch 1{
	case 1, 5, 10:
		fmt.Println("one, five or 10")
	case 2:
		fmt.Println("two")
	default: 
		fmt.Println("none one or two")
	}
}