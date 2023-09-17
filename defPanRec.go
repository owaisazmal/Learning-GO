package main

import (
	"fmt"
	"log"
)
// Panic := The panic() function in Go is used to signal that an unrecoverable error has occurred. 
func main(){
	fmt.Println("Start")
	panicker()
	fmt.Println("end")
}
func panicker(){

	fmt.Println("about to panic")
	defer func(){
		if err := recover(); err != nil{
			log.Println("Error: ", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("Done Panicking")

}