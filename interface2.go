package main

import "fmt"
//Incremener
func main() {
	myInt := IntCounter(0)
	var inc Incrementor = &myInt
	for i := 0; i < 10; i++{
		fmt.Println(inc.Increment())
	}
}

type Incrementor interface{
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int{
	*ic++
	return int(*ic)
}