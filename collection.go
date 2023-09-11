package main

import "fmt"
//a struct is a composite data type that groups together variables (fields) under a single name. 
//It is similar to a class or record in other programming languages and is used to represent a collection of related data fields. 
//Each field in a struct can have a different data type.
type Doctor struct {
	Num        int
	ActorName  string
	//episode    []string
	companions []string
}

func main() {
	aDoctor := Doctor{
		Num:       3,
		ActorName: "Jon Petre",
		companions: []string{
			"Usaibah",
			"Kokab",
			"Tariq",
		},
	}
	fmt.Println(aDoctor.ActorName)
	fmt.Println(aDoctor.companions)
	fmt.Println(aDoctor.companions[1])
}
