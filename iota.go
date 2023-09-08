package main

import "fmt"

const(
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main(){

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin?", isAdmin & roles == isAdmin)
	fmt.Printf("Is Admin?", isHeadquarters & roles == isHeadquarters)

}