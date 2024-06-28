package main

import (
	"fmt"
	"sync"
)

func hello() {
	fmt.Println("Hello Goroutine!")
}

var wg sync.WaitGroup

func hello2(i int) {
	defer wg.Done()
	fmt.Println(i)
}
func main() {
	// go hello()
	// fmt.Println("main goroutine done!")
	// time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello2(i)
	}
	wg.Wait()
}
