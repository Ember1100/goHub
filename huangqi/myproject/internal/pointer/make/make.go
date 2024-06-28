package main

import "fmt"

func main() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
}
