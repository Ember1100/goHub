package main

import "fmt"

func main() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}
