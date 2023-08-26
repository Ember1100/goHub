package main

import "fmt"

//package bool
func main() {
	//bool 布尔类型 只能赋值false或true
	var bool1 bool = false
	fmt.Println(bool1)
	var bool2 bool = 1 > 2
	fmt.Println(bool2)
	bool3 := 5 < 12
	fmt.Println(bool3)
}
