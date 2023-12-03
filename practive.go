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

//Citation:

/* 
https://youtu.be/_mVLhNgQ7_8?si=GhMlmHSpdpEq7Xzm

Slices explained by Geeks for Geeks

https://youtu.be/CrgD_otSzDI?si=N7TX5Fj6cwR-T1gY

https://youtu.be/lB5fMdUU8DE?si=QPpRUdqu-eT6U7iJ
*/

//****************Q3*******************

// import (
//     "fmt"
//     "math/rand"
// )

// func intRand2DArray(row, col int) [][]int {
    
// 	Arr := [][]int{}

//     for i := 0; i < row; i++ {

//         Row := []int{}

//         for j := 0; j < col; j++ {
//             Row = append(Row, rand.Intn(col))
//         }

//         Arr = append(Arr, Row)
//     }

//     return Arr
// }

// func main() {

//     row := 3 
//     col := 4 

//     intArr := intRand2DArray(row, col) 

//     for i := range intArr {
//         for j := range intArr[i] {
//             fmt.Printf("%d ", intArr[i][j])
//         }
//         fmt.Println()
//     }
// }

//****************Q4*******************

// import (
// 	"fmt"
// )

// func int2DArrayToString(input [][]int) string {

//     var res string

// 	for i := 0; i < len(input); i++ {

// 		row := input[i]
// 		var rowStrng string
		
//         for j := 0; j < len(row); j++ {
// 			rowStrng += fmt.Sprintf(" %6d", row[j])
// 		}
	
//         res += rowStrng + "\n"
	
//     }

// 	return res
// }
// //Citation
// //https://stackoverflow.com/questions/72191437/how-do-i-print-a-2-dimensional-array-as-a-grid-in-golang

// func main() {
// 	// Example usage
// 	arr := [][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}

// 	output := int2DArrayToString(arr)
// 	fmt.Println(output)
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

// import "fmt"


// func orderedArray(size int) []int {
    
    
    // arr := make([]int, size)
    
    
	// for i := 0; i < size; i++ {
        // arr[i] = i
		// }
        
        
        // return arr
        // }
        
        
        // func main() {
            
            // res := orderedArray(10)
            // fmt.Println(res)
            
            // }
            
//****************Q7*******************

// import "fmt"

// func ordered2DArray(row, col int) [][]int {
	
    //      var res [][]int
    // 	 for i := 0; i < row; i++ {
        
        // 		var rowVal []int
        
        // 		 for j := i * 10; j < i*10+col; j++ {
            // 			rowVal = append(rowVal, j)
            // 		}
            
            // 		 res = append(res, rowVal)
            
            //     }
            
            // 	return res
            // }
            
            // func main() {
	
                // 	row := 2
// 	col := 2
// 	arr := ordered2DArray(row, col)

// 		for _, row := range arr {
    // 		fmt.Println(row)
    // 	}
    // }

//****************Q8*******************

// import "fmt"

// func sumArray(input []int) int 
// {
    
    //     var sum int = 0
    
    //     for i := 0; i < len(input); i++ {
        //         val := input[i]
//         sum += val   
//     }

//     return sum

// }

// func main() {
    //     // Example usage
    //     arr := []int{1, 2, 3, 4, 5}
//     result := sumArray(arr)
//     fmt.Println(result)
// }


//****************Q9*******************

// import "fmt"

// func averageArray(input []int) int {
    
//     var total int = 0
    
//     for i := 0; i < len(input); i++ {
        
//         value := input[i]
//         total += value
//     }
    
//     return total / len(input)
// }

// func main() {
//     // Example usage
//     arr := []int{1, 2, 3, 4, 5}
//     result := averageArray(arr)
//     fmt.Println(result)
// }

//****************Q10*******************

import "fmt"

func sum2DArray(input [][]int) int {

    total := 0
    
    for i := 0; i < len(input); i++ {
        row := input[i]
        rowTotal := 0
    
        for j := 0; j < len(row); j++ {
            val := row[j]
            rowTotal += val
        }
    
        total += rowTotal
    
    }
    
    return total
}

func main() {
    // Example usage
    arr := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    result := sum2DArray(arr)
    fmt.Println(result)
}
