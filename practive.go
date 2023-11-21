package main

//****************Q1*******************

// func Print(input string){
// 	fmt.Println(input)
// }

//****************Q2*******************

// import (
//   "fmt"
//   "math/rand"
// )

// func main() {

//   randArr := intRandom1DArray(2)

//     fmt.Println(randArr)
// }

// func intRandom1DArray(size int) ([]int) {

//   res := make([]int, size)

//     for i := 0; i < size; i++{

//       res[i] = rand.Int()

//     }

//   return res
// }

//****************Q5*******************

// import "fmt"

// func main() {

// 	name := promptForInput("What's your name Homie")
// 	fmt.Println("Sup Homie ", name, "!")

// }

// func promptForInput(msg string) string {

// 	fmt.Print(msg + ": ")
// 	var input string
// 	fmt.Scanln(&input)

// 	return input

// }

//Citation
//https://www.w3schools.com/go/go_input.php


//****************Q6*******************

//NOTE: ONLY WORKS IN ECLIPSE BECAUSE OF GOLANG SPECIFICATIONS
//go get -u github.com/mohae/deepcopy

import (
	"fmt"
	"github.com/pkg/go-stream" 
  )
  
  func main() {
	require (
	  github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 
	  github.com/pkg/go-stream v1.0.0
	)
	orderedArray(10)
  }
  
  func orderedArray(size int) []int {
	return gostream.IntStream(0, size).ToArray()
  }
  
