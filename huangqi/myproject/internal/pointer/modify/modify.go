package main

import "fmt"

func main() {
	a := 10
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
