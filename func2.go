package main

import "fmt"

func main(){
	sayGreetings("Hello", "Owais")
}
func sayGreetings(greeting, name string){
	fmt.Println(greeting, name)
}