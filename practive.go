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



import "fmt"

func orderedArray(size int) []int {

    arr := make([]int, size)

    for i := 0; i < size; i++ {
        arr[i] = i
    }

    return arr
}

func main() {

    res := orderedArray(10)
    fmt.Println(res) 

}
