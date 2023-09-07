package main

import "fmt"

func main(){
	//Complex Numbers ---------- Majorly used for data science

	//var n complex128 = 1 + 2i								
	var n complex128 = complex(1, 2)    		// same thing
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))


}