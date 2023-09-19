package main

import "fmt"

func main() {
	g := greeter{
		greeting : "Hello",
		name: "Go",
	}
	g.greet()		//method invocation
}
type greeter struct{		value creator
	greeting string
	name string
}
// type counter int{
	
// }	
func (g greeter)greet(){		//value reciever
	fmt.Println(g.greeting, g.name)
}
// func (g *greeter)greet(){		//pointer reciever
// 	fmt.Println(g.greeting, g.name)
// }