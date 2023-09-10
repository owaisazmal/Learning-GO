package main

import "fmt"

//Maps are used to store data values in key:value pairs. Each element in a map is a key:value pair.
func main() {

	//StatePopulations := make(map[string]int)

	StatePopulations := map[string]int{ //Common syntax
		"California": 65843654,
		"Texas":      546354635,
		"New York":   546546,
		"ohio":       5541465,
	}

	//m := map[[3]int]string{}
	StatePopulations["ohio"] = 5323213 //can add values
	delete(StatePopulations, "Texas")  //delete
	fmt.Println(StatePopulations)
	fmt.Println(len(StatePopulations))
	fmt.Println(StatePopulations["ohio"])
}
