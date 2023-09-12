package main

import "fmt"

func main(){
	StatePopulations := map[string]int{ 
		"California": 65843654,
		"Texas":      546354635,
		"New York":   546546,
		"ohio":       5541465,
}

		if pop, ok := StatePopulations["Texas"]; ok{
			fmt.Println(pop)
		}
}
