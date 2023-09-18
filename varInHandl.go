package main

import "fmt"

func main(){
	a := []int{1,2,3,4}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)

	x := map[string]string{"foo": "car", "baz": "Buz"}
	y := x
	fmt.Println(x, y)
	x["foo"] = "qux"
	fmt.Println(x, y)
}