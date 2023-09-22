// 22 go routines
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter = 0
var counterMutex sync.Mutex

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	counterMutex.Lock()
	defer counterMutex.Unlock()
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counterMutex.Lock()
	defer counterMutex.Unlock()
	counter++
	wg.Done()
}
