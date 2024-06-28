package main

import "fmt"

const (
	n1 = iota //0
	n2        //1
	_
	_
	n4 //3
)

func main() {
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n4)

}
