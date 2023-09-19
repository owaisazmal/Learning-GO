package main

import "fmt"
//Annonymous Function
func main() {
	func(){											//immediately invoked function
		msg := "Hello GO!"							//isolated scope	
		fmt.Println(msg)
	}()
		// for i := 0; i < 5; i++ {			//syncrynous 
		//  func (i int){
		// 	fmt.Println(i)

		//  	}(i)
		// }
	}