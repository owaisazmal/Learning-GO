package main

import "fmt"

func main(){
	// s := []int{1, 2, 3}

	// for k, v := range s {			//key and value		(remove v if you don't want values)
	// 	fmt.Println(k, v)
	// }
	

	StatePopulations := map[string]int{ //Common syntax
		"California": 65843654,
		"Texas":      546354635,
		"New York":   546546,
		"ohio":       5541465,
	}
	for k, v := range StatePopulations {			//key and value
		fmt.Println(k, v)
	}
	for _, v := range StatePopulations {			//value
		fmt.Println(v)
	}
	for k := range StatePopulations {			//Key
		fmt.Println(k)
	}

	// x := "Hello"
	// for k, v := range x {			//key and value
	// 	fmt.Println(k, v)
	// }
}