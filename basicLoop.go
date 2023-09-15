package main

import "fmt"

func main(){
	for i, j := 0, 0; i < 5; i, j = i+1, j+1{		//can't use i++ and j++ together
		fmt.Println(i, j)
	}
	j := 0							//cleaner way
	for j < 5 {
		fmt.Print(j)
		j++
	}
}