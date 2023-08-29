package main

import "fmt"

//package switch
func main() {
	var num int = 1
	switch num {
	case 1, 0:
		fmt.Println(1)
	default:
		fmt.Println("不是数字1")
	}
}
