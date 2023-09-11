package main

import "fmt"

type Doctor struct {
    num        int
    actorName  string
    companions []string
}

func main() {
    aDoctor := Doctor{
        num:        3,
        actorName:  "Jon Petre",
        companions: []string{
            "Usaibah",
            "Kokab",
            "Tariq",
        },
    }
    fmt.Println(aDoctor.actorName)
    fmt.Println(aDoctor.companions)
}
