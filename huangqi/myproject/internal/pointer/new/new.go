package main

import "fmt"

func main() {
	//var a *int
	//*a = 100
	//fmt.Println(a)
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}
