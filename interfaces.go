package main

import "fmt"
//basics
//structs are data coontainer
//interface describes behaviour
func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello GO!"))
}

type Writer interface{
	Write([]byte)(int, error)
}
type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error){
	n, err := fmt.Println(string(data))
	return n, err
}
