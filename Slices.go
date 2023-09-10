package main

import "fmt"

func main() {
	
	a := []int{1,2,3}
	b := a	
	b[1] = 9

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("length: %v\n", len(a))				//length and capacity are different
	fmt.Printf("Capacity: %v\n", cap(a))
	

}