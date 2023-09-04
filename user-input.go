package main

import "fmt"

func main() {
	var fName string
	fmt.Println("Enter First Name: ")
	fmt.Scanln(&fName)

	var lName string
	fmt.Println("Enter Last Name: ")
	fmt.Scan(&lName)

	fmt.Print("Full Name is: " + fName + " " + lName)
}
