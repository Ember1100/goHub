package main

import "fmt"

func hello() {
	fmt.Println("Hello Goroutine!")
}
func main() {
	hello()
	fmt.Println("main goroutine done!")
	go hello()
}
