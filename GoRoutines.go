package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{

}
func main() {
	var msg = "Hello"
	go func (msg string)  {
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	//go sayHello()
	time.Sleep(100 * time.Millisecond)			//bad practice
}

// func sayHello(){
// 	fmt.Println("Hello")
// }