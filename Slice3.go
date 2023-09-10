package main

import "fmt"

//Built in make and append feature
func main() {

	a := make([]int, 3, 100) //(type, length, Capacity)~make feature
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	fmt.Println("\n\n\n")

	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	fmt.Println("\n\n\n")

	a = append(a, 2, 3, 4, 5) //can take two of more values
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

}
