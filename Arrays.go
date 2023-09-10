package main

import "fmt"

func main() {

	grades := [5]int{97, 74, 23}
	fmt.Printf("Grades: %v", grades)

	var Students [3]string
	fmt.Printf("Grades: %v\n", Students)
	Students[0] = "Owais"
	Students[1] = "Ayaan"
	Students[2] = "Ahmed"
	fmt.Printf("Students #1: %v\n", Students[1])
	fmt.Printf("Number of Students: %v\n", len(Students))
}
