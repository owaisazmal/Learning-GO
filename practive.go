package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

}

func Print(input string){
	fmt.Println(input)
}

// func Print(input int){
// 	fmt.Println(input)
// }

// func Print(input double){
// 	fmt.Println(input)
// }

func intRandom1DArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(100) // Adjust 100 to the desired range for random integers
	}

	return array
}